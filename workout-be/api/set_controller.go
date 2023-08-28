package api

import (
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSetRequest struct {
	ExerciseName string `json:"exercise_name" binding:"required"`
	SetNumber    int32  `json:"set_number" binding:"required"`
	Weight       int32  `json:"weight" binding:"required"`
	RestDuration string `json:"rest_duration" binding:"required"`
	Notes        string `json:"notes" binding:"required"`
}

type getSetRequest struct {
	SetID int64 `uri:"id" binding:"required,min=1"`
}

type updateSetRequest struct {
	SetID        int64  `json:"set_id"`
	SetNumber    int32  `json:"set_number"`
	Weight       int32  `json:"weight"`
	RestDuration string `json:"rest_duration"`
	Notes        string `json:"notes"`
}

func (server *Server) createSet(ctx *gin.Context) {
	var req createSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateSetParams{
		ExerciseName: req.ExerciseName,
		SetNumber:    req.SetNumber,
		Weight:       req.Weight,
		RestDuration: req.RestDuration,
		Notes:        req.Notes,
	}
	set, err := server.store.CreateSet(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, set)
}

func (server *Server) getSet(ctx *gin.Context) {
	var req getSetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	set, err := server.store.GetSet(ctx, req.SetID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Check if the profile being updated belongs to the authenticated user
	ctx.JSON(http.StatusOK, set)
}

func (server *Server) updateSet(ctx *gin.Context) {
	var req updateSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateSetParams{
		SetID:        req.SetID,
		SetNumber:    req.SetNumber,
		Weight:       req.Weight,
		RestDuration: req.RestDuration,
		Notes:        req.Notes,
	}

	set, err := server.store.UpdateSet(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, set)
}
