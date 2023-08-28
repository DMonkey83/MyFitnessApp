package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/gin-gonic/gin"
)

type createExerciseRequest struct {
	EquipmentRequired util.Equipmenttype   `json:"equipment_required" binding:"required,equipment"`
	ExerciseName      string               `json:"exercise_name" binding:"required"`
	Description       string               `json:"description"`
	MuscleGroupName   util.Musclegroupenum `json:"muscle_group_name" binding:"required,mucle_group"`
}

type getExerciseRequest struct {
	ExerciseName string `uri:"exercise_name" binding:"required,min=1"`
}

type updateExerciseRequest struct {
	ExerciseID        int64                `json:"exercise_id" binding:"required,min=1"`
	EquipmentRequired util.Equipmenttype   `json:"equipment_required" binding:"equipment"`
	ExerciseName      string               `json:"exercise_name" binding:"required"`
	Description       string               `json:"description"`
	MuscleGroupName   util.Musclegroupenum `json:"muscle_group_name" binding:"mucle_group"`
}

type listAllExercisesRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=5,max=10"`
}

type listEquipmentExercisesRequest struct {
	Limit             int32              `form:"limit" binding:"required,min=1"`
	Offset            int32              `form:"offset" binding:"required,min=5,max=10"`
	EquipmentRequired util.Equipmenttype `form:"equipment_required" binding:"required,equipment"`
}

type listMuscleGroupExercisesRequest struct {
	Limit           int32                `form:"limit" binding:"required,min=1"`
	Offset          int32                `form:"offset" binding:"required,min=5,max=10"`
	MuscleGroupName util.Musclegroupenum `form:"muscle_group_name" binding:"mucle_group"`
}

func (server *Server) createExercise(ctx *gin.Context) {
	var req createExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateExerciseParams{
		ExerciseName:      req.ExerciseName,
		Description:       req.Description,
		EquipmentRequired: db.Equipmenttype(req.EquipmentRequired),
		MuscleGroupName:   db.Musclegroupenum(req.MuscleGroupName),
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
		Description:       req.Description,
		EquipmentRequired: db.Equipmenttype(req.EquipmentRequired),
		MuscleGroupName:   db.Musclegroupenum(req.MuscleGroupName),
	}

	exercise, err := server.store.UpdateExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) listMuscleGroupExercises(ctx *gin.Context) {
	var req listMuscleGroupExercisesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.ListMuscleGroupExercisesParams{
		Limit:           req.Limit,
		Offset:          (req.Offset - 1) * req.Limit,
		MuscleGroupName: db.Musclegroupenum(req.MuscleGroupName),
	}
	exercises, err := server.store.ListMuscleGroupExercises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}

func (server *Server) listEquipmentExercises(ctx *gin.Context) {
	var req listEquipmentExercisesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.ListEquipmentExercisesParams{
		Limit:             req.Limit,
		Offset:            (req.Offset - 1) * req.Limit,
		EquipmentRequired: db.Equipmenttype(req.EquipmentRequired),
	}
	exercises, err := server.store.ListEquipmentExercises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}

func (server *Server) listAllExercises(ctx *gin.Context) {
	var req listAllExercisesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.ListAllExercisesParams{
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
	}
	exercises, err := server.store.ListAllExercises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}
