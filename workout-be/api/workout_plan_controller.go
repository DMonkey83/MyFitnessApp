package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createPlanRequest struct {
	Username    string             `json:"username"`
	PlanName    string             `json:"plan_name"`
	Description string             `json:"description"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     time.Time          `json:"end_date"`
	Goal        db.Workoutgoalenum `json:"goal" binding:"required,oneof='BuildMuscle' 'Lose Weight' 'Improve Endurance' 'Maintain Fitness' 'Tone Body' 'Custom'"`
	Difficulty  db.Difficulty      `json:"difficulty" binding:"required,oneof='Very Light' 'Light' 'Moderate' 'Heavy' 'Very Heavy'"`
	IsPublic    db.Visibility      `json:"is_public" binding:"required,oneof=Private Public"`
}

type getPlanRequest struct {
	PlanID   int64  `uri:"plan_id"`
	Username string `uri:"username"`
}

type updatePlanRequest struct {
	PlanID      int64              `json:"plan_id"`
	Username    string             `json:"username"`
	PlanName    string             `json:"plan_name"`
	Description string             `json:"description"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     time.Time          `json:"end_date"`
	Goal        db.Workoutgoalenum `json:"goal" binding:"required,oneof='BuildMuscle' 'Lose Weight' 'Improve Endurance' 'Maintain Fitness' 'Tone Body' 'Custom'"`
	Difficulty  db.Difficulty      `json:"difficulty" binding:"required,oneof='Very Light' 'Light' 'Moderate' 'Heavy' 'Very Heavy'"`
	IsPublic    db.Visibility      `json:"is_public" binding:"required,oneof=Private Public"`
}

func (server *Server) createPlan(ctx *gin.Context) {
	var req createPlanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreatePlanParams{
		Username:    authPayload.Username,
		PlanName:    req.PlanName,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Goal:        req.Goal,
		Difficulty:  req.Difficulty,
		IsPublic:    req.IsPublic,
	}
	plan, err := server.store.CreatePlan(ctx, arg)

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

func (server *Server) getPlan(ctx *gin.Context) {
	var req getPlanRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.GetPlanParams{
		PlanID:   req.PlanID,
		Username: req.Username,
	}

	plan, err := server.store.GetPlan(ctx, arg)
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
	if plan.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, plan)
}

func (server *Server) updatePlan(ctx *gin.Context) {
	var req updatePlanRequest
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

	arg := db.UpdatePlanParams{
		PlanID:      req.PlanID,
		Username:    authPayload.Username,
		PlanName:    req.PlanName,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Goal:        req.Goal,
		Difficulty:  req.Difficulty,
		IsPublic:    req.IsPublic,
	}

	userProfile, err := server.store.UpdatePlan(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}
