package api

import (
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserProfileRequest struct {
	Username      string        `json:"username" binding:"required"`
	FullName      string        `json:"full_name" binding:"required"`
	Age           int32         `json:"age" binding:"required"`
	Gender        string        `json:"gender" binding:"required,oneof=female male"`
	HeightCm      float64       `json:"height_cm" binding:"required"`
	HeightFtIn    pgtype.Text   `json:"height_ft_in"`
	PreferredUnit db.Weightunit `json:"preferred_unit" binding:"required"`
}

type getUserProfileRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
}

func (server *Server) createUserProfile(ctx *gin.Context) {
	var req createUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateUserProfileParams{
		Username:      req.Username,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: req.PreferredUnit,
	}
	userProfile, err := server.store.CreateUserProfile(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}

func (server *Server) getUserProfile(ctx *gin.Context) {
	var req getUserProfileRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userProfile, err := server.store.GetUserProfile(ctx, req.Username)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}
