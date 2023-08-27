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

type createWeightEntryRequest struct {
	Username  string    `json:"username"`
	EntryDate time.Time `json:"entry_date"`
	WeightKg  int32     `json:"weight_kg"`
	WeightLb  int32     `json:"weight_lb"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

type getWeightEntryRequest struct {
	WeightEntryID int64 `uri:"entry_id"`
}

type updateWeightEntryRequest struct {
	WeightEntryID int64     `json:"weight_entry_id"`
	Username      string    `json:"username"`
	EntryDate     time.Time `json:"entry_date"`
	WeightKg      int32     `json:"weight_kg"`
	WeightLb      int32     `json:"weight_lb"`
	Notes         string    `json:"notes"`
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

	entry, err := server.store.GetWeightEntry(ctx, req.WeightEntryID)
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
