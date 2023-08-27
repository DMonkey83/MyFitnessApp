package api

import (
	"net/http"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createWorkoutLogRequest struct {
	Username            string          `json:"username"`
	PlanID              int64           `json:"plan_id"`
	LogDate             time.Time       `json:"log_date"`
	Rating              db.Rating       `json:"rating" binding:"required,oneof=1 2 3 4 5"`
	FatigueLevel        db.Fatiguelevel `json:"fatigue_level" binding:"required,oneof='Very Light' 'Light' 'Modelrate' 'Heavy' 'Very Heavy'"`
	OverallFeeling      string          `json:"overall_feeling"`
	Comments            string          `json:"comments"`
	WorkoutDuration     string          `json:"workout_duration"`
	TotalCaloriesBurned int32           `json:"total_calories_burned"`
	TotalDistance       int32           `json:"total_distance"`
	TotalRepetitions    int32           `json:"total_repetitions"`
	TotalSets           int32           `json:"total_sets"`
	TotalWeightLifted   int32           `json:"total_weight_lifted"`
}

type getWorkoutLogRequest struct {
	LogID int64 `uri:"log_id" binding:"required,min=1"`
}

type updateWorkoutLogRequest struct {
	LogID               int64           `json:"log_id"`
	LogDate             time.Time       `json:"log_date"`
	WorkoutDuration     string          `json:"workout_duration"`
	Comments            string          `json:"comments"`
	Rating              db.Rating       `json:"rating" binding:"required,oneof=1 2 3 4 5"`
	FatigueLevel        db.Fatiguelevel `json:"fatigue_level" binding:"required,oneof='Very Light' 'Light' 'Modelrate' 'Heavy' 'Very Heavy'"`
	TotalSets           int32           `json:"total_sets"`
	TotalDistance       int32           `json:"total_distance"`
	TotalRepetitions    int32           `json:"total_repetitions"`
	TotalWeightLifted   int32           `json:"total_weight_lifted"`
	TotalCaloriesBurned int32           `json:"total_calories_burned"`
	OverallFeeling      string          `json:"overall_feeling"`
}

func (server *Server) createWorkoutLog(ctx *gin.Context) {
	var req createWorkoutLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateWorkoutLogParams{
		Username:            authPayload.Username,
		PlanID:              req.PlanID,
		LogDate:             req.LogDate,
		Rating:              req.Rating,
		FatigueLevel:        req.FatigueLevel,
		OverallFeeling:      req.OverallFeeling,
		Comments:            req.Comments,
		WorkoutDuration:     req.WorkoutDuration,
		TotalCaloriesBurned: req.TotalCaloriesBurned,
		TotalDistance:       req.TotalDistance,
		TotalRepetitions:    req.TotalRepetitions,
		TotalSets:           req.TotalSets,
		TotalWeightLifted:   req.TotalWeightLifted,
	}
	log, err := server.store.CreateWorkoutLog(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, log)
}

func (server *Server) getWorkoutLog(ctx *gin.Context) {
	var req getWorkoutLogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	log, err := server.store.GetWorkoutLog(ctx, req.LogID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, log)
}

func (server *Server) updateWorkoutLog(ctx *gin.Context) {
	var req updateWorkoutLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateWorkoutLogParams{
		LogDate:             req.LogDate,
		Rating:              db.Rating1,
		FatigueLevel:        req.FatigueLevel,
		OverallFeeling:      req.OverallFeeling,
		Comments:            req.Comments,
		WorkoutDuration:     req.WorkoutDuration,
		TotalCaloriesBurned: req.TotalCaloriesBurned,
		TotalDistance:       req.TotalDistance,
		TotalRepetitions:    req.TotalRepetitions,
		TotalSets:           req.TotalSets,
		TotalWeightLifted:   req.TotalWeightLifted,
		LogID:               req.LogID,
	}

	userProfile, err := server.store.UpdateWorkoutLog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}
