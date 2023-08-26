package api

import (
	"errors"
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserProfileRequest struct {
	FullName      string        `json:"full_name" binding:"required"`
	Age           int32         `json:"age" binding:"required"`
	Gender        string        `json:"gender" binding:"required,oneof=female male"`
	HeightCm      float64       `json:"height_cm" binding:"required"`
	HeightFtIn    pgtype.Text   `json:"height_ft_in"`
	PreferredUnit db.Weightunit `json:"preferred_unit" binding:"required,oneof=kg lb"`
}

type getUserProfileRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
}

type updateUserProfileRequest struct {
	FullName      string        `json:"full_name"`
	Age           int32         `json:"age"`
	Gender        string        `json:"gender" binding:"oneof=female male"`
	HeightCm      float64       `json:"height_cm"`
	HeightFtIn    pgtype.Text   `json:"height_ft_in"`
	PreferredUnit db.Weightunit `json:"preferred_unit" binding:"oneof=kg lb"`
}

func (server *Server) createUserProfile(ctx *gin.Context) {
	var req createUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateUserProfileParams{
		Username:      authPayload.Username,
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

	// Check if the profile being updated belongs to the authenticated user
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if userProfile.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, userProfile)
}

func (server *Server) updateUserProfile(ctx *gin.Context) {
	var req updateUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != ctx.Param("username") {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	arg := db.UpdateUserProfileParams{
		Username:      authPayload.Username,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: req.PreferredUnit,
	}

	userProfile, err := server.store.UpdateUserProfile(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}
