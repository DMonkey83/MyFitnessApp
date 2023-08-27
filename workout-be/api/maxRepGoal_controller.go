package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createMaxRepRequest struct {
	Username     string `json:"username"`
	ExerciseName string `json:"exercise_name"`
	GoalReps     int32  `json:"goal_reps"`
	Notes        string `json:"notes"`
}

type getMaxRepRequest struct {
	ExerciseName string `uri:"exercise_name"`
	GoalID       int64  `uri:"goal_id"`
	Username     string `uri:"username"`
}

type updateMaxRepRequest struct {
	GoalID       int64  `json:"goal_id"`
	Username     string `json:"username"`
	ExerciseName string `json:"exercise_name"`
	GoalReps     int32  `json:"goal_reps"`
	Notes        string `json:"notes"`
}

func (server *Server) createMaxRepGoal(ctx *gin.Context) {
	var req createMaxRepRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateMaxRepGoalParams{
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
		GoalReps:     req.GoalReps,
		Notes:        req.Notes,
	}
	goal, err := server.store.CreateMaxRepGoal(ctx, arg)

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

func (server *Server) getMaxRepGoal(ctx *gin.Context) {
	var req getMaxRepRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.GetMaxRepGoalParams{
		ExerciseName: req.ExerciseName,
		Username:     req.Username,
		GoalID:       req.GoalID,
	}

	goal, err := server.store.GetMaxRepGoal(ctx, arg)
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

func (server *Server) updateMaxRepGoal(ctx *gin.Context) {
	var req updateMaxRepRequest
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

	arg := db.UpdateMaxRepGoalParams{
		Username:     authPayload.Username,
		GoalID:       req.GoalID,
		Notes:        req.Notes,
		GoalReps:     req.GoalReps,
		ExerciseName: req.ExerciseName,
	}

	goal, err := server.store.UpdateMaxRepGoal(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, goal)
}
