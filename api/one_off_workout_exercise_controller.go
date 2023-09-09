package api

import (
	"errors"
	"log"
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/token"
	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/gin-gonic/gin"
)

type createOneOffExerciseLogRequest struct {
	WorkoutID       int64                `json:"workout_id" binding:"required"`
	ExerciseName    string               `json:"exercise_name" binding:"required"`
	Description     string               `json:"description" binding:"required"`
	MuscleGroupName util.Musclegroupenum `json:"muscle_group_name" binding:"required,muscle_group"`
}

type getOneOffExerciseLogRequest struct {
	ID        int64  `uri:"id" binding:"required,min=1"`
	WorkoutID int64  `uri:"workout_id"`
	Username  string `uri:"username" binding:"required,min=1"`
}

type updateOneOffExerciseLogRequest struct {
	ID              int32                `json:"id"`
	WorkoutID       int64                `json:"workout_id"`
	Description     string               `json:"description"`
	MuscleGroupName util.Musclegroupenum `json:"muscle_group_name" binding:"required,muscle_group"`
	Username        string               `json:"username" binding:"required,min=1"`
}

type listOneOffExerciseLogRequests struct {
	Limit        int32  `form:"limit" binding:"required,min=1"`
	Offset       int32  `form:"offset" binding:"required,min=5,max=10"`
	ExerciseName string `form:"exercise_name" binding:"required"`
	Username     string `form:"username" binding:"required,min=1"`
	WorkoutID    int64  `form:"workout_id" binding:"required"`
}

func (server *Server) createOneOffExerciseLog(ctx *gin.Context) {
	var req createOneOffExerciseLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateOneOffWorkoutExerciseParams{
		WorkoutID:       req.WorkoutID,
		ExerciseName:    req.ExerciseName,
		Description:     req.Description,
		MuscleGroupName: db.Musclegroupenum(req.MuscleGroupName),
	}
	exercise, err := server.store.CreateOneOffWorkoutExercise(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) getOneOffExerciseLog(ctx *gin.Context) {
	var req getOneOffExerciseLogRequest
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

	arg := db.GetOneOffWorkoutExerciseParams{
		ID:        int32(req.ID),
		WorkoutID: req.WorkoutID,
	}

	exercise, err := server.store.GetOneOffWorkoutExercise(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Check if the profile being updated belongs to the authenticated user
	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) updateOneOffExerciseLog(ctx *gin.Context) {
	var req updateOneOffExerciseLogRequest
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

	arg := db.UpdateOneOffWorkoutExerciseParams{
		ID:              req.ID,
		WorkoutID:       req.WorkoutID,
		Description:     req.Description,
		MuscleGroupName: db.Musclegroupenum(req.MuscleGroupName),
	}

	exercise, err := server.store.UpdateOneOffWorkoutExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) listOneOffExerciseLogs(ctx *gin.Context) {
	var req listOneOffExerciseLogRequests
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
	arg := db.ListAllOneOffWorkoutExercisesParams{
		Limit:     req.Limit,
		Offset:    (req.Offset - 1) * req.Limit,
		WorkoutID: req.WorkoutID,
	}
	exercises, err := server.store.ListAllOneOffWorkoutExercises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}
