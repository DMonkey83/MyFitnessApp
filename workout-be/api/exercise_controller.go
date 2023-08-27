package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createExerciseRequest struct {
	EquipmentRequired db.Equipmenttype   `json:"equipment_required" binding:"required,oneof=Barbell Dumbbell Machine Bodyweight Other"`
	ExerciseName      string             `json:"exercise_name" binding:"required"`
	Description       string             `json:"description"`
	MuscleGroupName   db.Musclegroupenum `json:"muscle_group_name" binding:"required,oneof=Chest Back Legs Shoulders Arms Abs Cardio"`
}

type getExerciseRequest struct {
	ExerciseName string `uri:"exercise_name" binding:"required,min=1"`
}

type updateExerciseRequest struct {
	ExerciseID        int64              `json:"exercise_id" binding:"required,min=1"`
	EquipmentRequired db.Equipmenttype   `json:"equipment_required" binding:"required,oneof=Barbell Dumbbell Machine Bodyweight Other"`
	ExerciseName      string             `json:"exercise_name" binding:"required"`
	Description       string             `json:"description"`
	MuscleGroupName   db.Musclegroupenum `json:"muscle_group_name" binding:"required,oneof=Chest Back Legs Shoulders Arms Abs Cardio"`
}

func (server *Server) createExercise(ctx *gin.Context) {
	var req createExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateExerciseParams{
		ExerciseName:      req.ExerciseName,
		EquipmentRequired: req.EquipmentRequired,
		Description:       req.Description,
		MuscleGroupName:   req.MuscleGroupName,
	}
	exercise, err := server.store.CreateExercise(ctx, arg)

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

func (server *Server) getExercise(ctx *gin.Context) {
	var req getExerciseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	exercise, err := server.store.GetExercise(ctx, req.ExerciseName)
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

func (server *Server) updateExercise(ctx *gin.Context) {
	var req updateExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	log.Print(req)
	arg := db.UpdateExerciseParams{
		ExerciseName:      req.ExerciseName,
		EquipmentRequired: req.EquipmentRequired,
		Description:       req.Description,
		MuscleGroupName:   req.MuscleGroupName,
	}

	exercise, err := server.store.UpdateExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
