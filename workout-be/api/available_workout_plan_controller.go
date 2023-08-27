package api

import (
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createAvailablePlanRequest struct {
	PlanName        string                 `json:"plan_name"`
	Description     string                 `json:"description"`
	Goal            db.NullWorkoutgoalenum `json:"goal" binding:"required,oneof='BuildMuscle' 'Lose Weight' 'Improve Endurance' 'Maintain Fitness' 'Tone Body' 'Custom'"`
	Difficulty      db.NullDifficulty      `json:"difficulty" binding:"required,oneof='Very Light' 'Light' 'Moderate' 'Heavy' 'Very Heavy'"`
	IsPublic        db.NullVisibility      `json:"is_public" binding:"required,oneof=Private Public"`
	CreatorUsername string                 `json:"creator_username"`
}

type getAvailablePlanRequest struct {
	PlanID int64 `uri:"plan_id"`
}

type updateAvailablePlanRequest struct {
	PlanID      int64                  `json:"plan_id"`
	Description string                 `json:"description"`
	PlanName    string                 `json:"plan_name"`
	Goal        db.NullWorkoutgoalenum `json:"goal" binding:"required,oneof='BuildMuscle' 'Lose Weight' 'Improve Endurance' 'Maintain Fitness' 'Tone Body' 'Custom'"`
	Difficulty  db.NullDifficulty      `json:"difficulty" binding:"required,oneof='Very Light' 'Light' 'Moderate' 'Heavy' 'Very Heavy'"`
	IsPublic    db.NullVisibility      `json:"is_public" binding:"required,oneof=Private Public"`
}

func (server *Server) createAvailablePlan(ctx *gin.Context) {
	var req createAvailablePlanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAvailablePlanParams{
		CreatorUsername: authPayload.Username,
		PlanName:        req.PlanName,
		Description:     req.Description,
		Goal:            req.Goal,
		Difficulty:      req.Difficulty,
		IsPublic:        req.IsPublic,
	}
	plan, err := server.store.CreateAvailablePlan(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (server *Server) getAvailablePlan(ctx *gin.Context) {
	var req getAvailablePlanRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	plan, err := server.store.GetAvailablePlan(ctx, req.PlanID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, plan)
}

func (server *Server) updateAvailablePlan(ctx *gin.Context) {
	var req updateAvailablePlanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateAvailablePlanParams{
		PlanID:      req.PlanID,
		PlanName:    req.PlanName,
		Description: req.Description,
		Goal:        req.Goal,
		Difficulty:  req.Difficulty,
		IsPublic:    req.IsPublic,
	}

	userProfile, err := server.store.UpdateAvailablePlan(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}
