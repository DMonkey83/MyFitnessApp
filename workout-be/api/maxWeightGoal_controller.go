package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createMaxWeightRequest struct {
	Username     string `json:"username"`
	ExerciseName string `json:"exercise_name"`
	GoalWeight   int32  `json:"goal_weight"`
	Notes        string `json:"notes"`
}

type getMaxWeightRequest struct {
	ExerciseName string `uri:"exercise_name"`
	Username     string `uri:"username"`
	GoalID       int64  `uri:"goal_id"`
}

type updateMaxWeightRequest struct {
	GoalID       int64  `json:"goal_id"`
	Username     string `json:"username"`
	ExerciseName string `json:"exercise_name"`
	GoalWeight   int32  `json:"goal_weight"`
	Notes        string `json:"notes"`
}

func (server *Server) createMaxWeightGoal(ctx *gin.Context) {
	var req createMaxWeightRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateMaxWeightGoalParams{
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
		GoalWeight:   req.GoalWeight,
		Notes:        req.Notes,
	}
	goal, err := server.store.CreateMaxWeightGoal(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goal)
}

func (server *Server) getMaxWeightGoal(ctx *gin.Context) {
	var req getMaxWeightRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.GetMaxWeightGoalParams{
		ExerciseName: req.ExerciseName,
		Username:     req.Username,
		GoalID:       req.GoalID,
	}

	goal, err := server.store.GetMaxWeightGoal(ctx, arg)
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
	if goal.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, goal)
}

func (server *Server) updateMaxWeightGoal(ctx *gin.Context) {
	var req updateMaxWeightRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != req.Username {
		log.Println("auth", authPayload.Username, req.Username)
		log.Println("req", ctx.Param("username"))
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	arg := db.UpdateMaxWeightGoalParams{
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
		Notes:        req.Notes,
		GoalWeight:   req.GoalWeight,
	}

	entry, err := server.store.UpdateMaxWeightGoal(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}
