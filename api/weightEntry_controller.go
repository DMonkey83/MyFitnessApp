package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/token"
	"github.com/gin-gonic/gin"
)

type createWeightEntryRequest struct {
	Username  string    `json:"username" binding:"required"`
	EntryDate time.Time `json:"entry_date" binding:"required"`
	WeightKg  int32     `json:"weight_kg" binding:"required"`
	WeightLb  int32     `json:"weight_lb" binding:"required"`
	Notes     string    `json:"notes" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type getWeightEntryRequest struct {
	WeightEntryID int64  `uri:"entry_id" binding:"required"`
	Username      string `uri:"username" binding:"required"`
}

type updateWeightEntryRequest struct {
	WeightEntryID int64     `json:"weight_entry_id"`
	Username      string    `json:"username"`
	EntryDate     time.Time `json:"entry_date"`
	WeightKg      int32     `json:"weight_kg"`
	WeightLb      int32     `json:"weight_lb"`
	Notes         string    `json:"notes"`
}

type listWeightEntriesRequest struct {
	Limit    int32  `form:"limit" binding:"required,min=1"`
	Offset   int32  `form:"offset" binding:"required,min=5,max=10"`
	Username string `form:"username" binding:"required,min=1"`
}

func (server *Server) createWeightEntry(ctx *gin.Context) {
	var req createWeightEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateWeightEntryParams{
		Username:  authPayload.Username,
		EntryDate: req.EntryDate,
		WeightKg:  req.WeightKg,
		WeightLb:  req.WeightLb,
		Notes:     req.Notes,
	}
	entry, err := server.store.CreateWeightEntry(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

func (server *Server) getWeightEntry(ctx *gin.Context) {
	var req getWeightEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	arg := db.GetWeightEntryParams{
		WeightEntryID: req.WeightEntryID,
		Username:      req.Username,
	}

	entry, err := server.store.GetWeightEntry(ctx, arg)
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
	if entry.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, entry)
}

func (server *Server) updateWeightEntry(ctx *gin.Context) {
	var req updateWeightEntryRequest
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

	arg := db.UpdateWeightEntryParams{
		Username:      authPayload.Username,
		WeightEntryID: req.WeightEntryID,
		EntryDate:     req.EntryDate,
		WeightKg:      req.WeightKg,
		WeightLb:      req.WeightLb,
		Notes:         req.Notes,
	}

	entry, err := server.store.UpdateWeightEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

func (server *Server) listWeightEntries(ctx *gin.Context) {
	var req listWeightEntriesRequest
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
	arg := db.ListWeightEntriesParams{
		Limit:    req.Limit,
		Offset:   (req.Offset - 1) * req.Limit,
		Username: authPayload.Username,
	}
	exercises, err := server.store.ListWeightEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}
