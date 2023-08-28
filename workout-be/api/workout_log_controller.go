package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/gin-gonic/gin"
)

type createWorkoutLogRequest struct {
	Username            string            `json:"username" binding:"required"`
	PlanID              int64             `json:"plan_id" binding:"required"`
	LogDate             time.Time         `json:"log_date" binding:"required"`
	Rating              util.Rating       `json:"rating" binding:"required,rating"`
	FatigueLevel        util.Fatiguelevel `json:"fatigue_level" binding:"required,fatigue_level"`
	OverallFeeling      string            `json:"overall_feeling" binding:"required"`
	Comments            string            `json:"comments" binding:"required"`
	WorkoutDuration     string            `json:"workout_duration" binding:"required"`
	TotalCaloriesBurned int32             `json:"total_calories_burned" binding:"required"`
	TotalDistance       int32             `json:"total_distance" binding:"required"`
	TotalRepetitions    int32             `json:"total_repetitions" binding:"required"`
	TotalSets           int32             `json:"total_sets" binding:"required"`
	TotalWeightLifted   int32             `json:"total_weight_lifted" binding:"required"`
}

type getWorkoutLogRequest struct {
	LogID    int64  `uri:"log_id" binding:"required,min=1"`
	Username string `uri:"username" binding:"required"`
}

type updateWorkoutLogRequest struct {
	Username            string            `json:"username" binding:"required"`
	LogID               int64             `json:"log_id"`
	LogDate             time.Time         `json:"log_date"`
	WorkoutDuration     string            `json:"workout_duration"`
	Comments            string            `json:"comments"`
	Rating              util.Rating       `json:"rating" binding:"rating"`
	FatigueLevel        util.Fatiguelevel `json:"fatigue_level" binding:"fatigue_level"`
	TotalSets           int32             `json:"total_sets"`
	TotalDistance       int32             `json:"total_distance"`
	TotalRepetitions    int32             `json:"total_repetitions"`
	TotalWeightLifted   int32             `json:"total_weight_lifted"`
	TotalCaloriesBurned int32             `json:"total_calories_burned"`
	OverallFeeling      string            `json:"overall_feeling"`
}

type listWorkoutLogsRequest struct {
	Limit    int32  `form:"limit" binding:"required,min=1"`
	Offset   int32  `form:"offset" binding:"required,min=5,max=10"`
	Username string `form:"username" binding:"required,min=1"`
	PlanID   int64  `form:"plan_id" binding:"required,min=1"`
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
		Rating:              db.Rating(req.Rating),
		FatigueLevel:        db.Fatiguelevel(req.FatigueLevel),
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != req.Username {
		log.Println("auth", authPayload.Username, req.Username)
		log.Println("req", ctx.Param("username"))
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != req.Username {
		log.Println("auth", authPayload.Username, req.Username)
		log.Println("req", ctx.Param("username"))
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	arg := db.UpdateWorkoutLogParams{
		LogDate:             req.LogDate,
		Rating:              db.Rating1,
		FatigueLevel:        db.Fatiguelevel(req.FatigueLevel),
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

func (server *Server) listWorktoutLogs(ctx *gin.Context) {
	var req listWorkoutLogsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
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
	arg := db.ListWorkoutLogsParams{
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
		PlanID: req.PlanID,
	}
	exercises, err := server.store.ListWorkoutLogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}
