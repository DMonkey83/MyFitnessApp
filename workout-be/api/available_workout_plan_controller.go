package api

import (
	"net/http"

	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/gin-gonic/gin"
)

type createAvailablePlanRequest struct {
	PlanName        string               `json:"plan_name" binding:"required"`
	Description     string               `json:"description" binding:"required"`
	Goal            util.Workoutgoalenum `json:"goal" binding:"required,goal"`
	Difficulty      util.Difficulty      `json:"difficulty" binding:"required,difficulty"`
	IsPublic        util.Visibility      `json:"is_public" binding:"required,is_public"`
	CreatorUsername string               `json:"creator_username" binding:"required"`
}

type getAvailablePlanRequest struct {
	PlanID int64 `uri:"plan_id"`
}

type updateAvailablePlanRequest struct {
	PlanID      int64                `json:"plan_id"`
	Description string               `json:"description"`
	PlanName    string               `json:"plan_name"`
	Goal        util.Workoutgoalenum `json:"goal" binding:"required,goal"`
	Difficulty  util.Difficulty      `json:"difficulty" binding:"required,difficulty"`
	IsPublic    util.Visibility      `json:"is_public" binding:"required,is_public"`
}

type listAllAvailablePlansRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required"`
}

type listAllAvailablePlansByCreatorRequest struct {
	Limit           int32  `form:"limit" binding:"required,min=5,max=20"`
	Offset          int32  `form:"offset" binding:"required,min=1"`
	CreatorUsername string `form:"creator" binding:"required"`
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
		Goal:            db.Workoutgoalenum(req.Goal),
		Difficulty:      db.Difficulty(req.Difficulty),
		IsPublic:        db.Visibility(req.IsPublic),
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
		Goal:        db.Workoutgoalenum(req.Goal),
		Difficulty:  db.Difficulty(req.Difficulty),
		IsPublic:    db.Visibility(req.IsPublic),
	}

	userProfile, err := server.store.UpdateAvailablePlan(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userProfile)
}

func (server *Server) listAllAvailablePlansByCreator(ctx *gin.Context) {
	var req listAllAvailablePlansByCreatorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.ListAvailablePlansByCreatorParams{
		Limit:           req.Limit,
		Offset:          (req.Offset - 1) * req.Limit,
		CreatorUsername: req.CreatorUsername,
	}
	plans, err := server.store.ListAvailablePlansByCreator(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, plans)
}

func (server *Server) listAllAvailablePlans(ctx *gin.Context) {
	var req listAllAvailablePlansRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.ListAllAvailablePlansParams{
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
	}
	plans, err := server.store.ListAllAvailablePlans(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, plans)
}
