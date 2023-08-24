package api

import (
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserProfileRequest struct {
	UserID        int64         `json:"user_id" binding:"required"`
	FullName      string        `json:"full_name" binding:"required"`
	Age           int32         `json:"age" binding:"required"`
	Gender        string        `json:"gender" binding:"required,oneof=female male"`
	HeightCm      float64       `json:"height_cm" binding:"required"`
	HeightFtIn    pgtype.Text   `json:"height_ft_in"`
	PreferredUnit db.Weightunit `json:"preferred_unit" binding:"required"`
}

type getUserProfileRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createUserProfile(ctx *gin.Context) {
	var req createUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateUserProfileParams{
		UserID:        req.UserID,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: req.PreferredUnit,
	}
	userProfile, err := server.store.CreateUserProfile(ctx, arg)

	if err != nil {
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

	userProfile, err := server.store.GetUserProfile(ctx, req.ID)
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
