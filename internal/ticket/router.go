package ticket

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Ridju/ticktr/config"
	db "github.com/Ridju/ticktr/internal/db/sqlc"
	"github.com/Ridju/ticktr/internal/middleware"
	"github.com/Ridju/ticktr/internal/token"
	"github.com/gin-gonic/gin"
)

type ticketRouter struct {
	repo   ITicketRepository
	ts     ITicketService
	config config.Config
}

func InitTicketRouter(r gin.IRoutes, store db.Store, config config.Config) {
	ticketRepo := NewTicketRepository(store)
	ticketService := NewTicketService(ticketRepo)

	ticketRouter := ticketRouter{
		repo:   ticketRepo,
		ts:     ticketService,
		config: config,
	}

	r.POST("/ticket", ticketRouter.createTicket)
	r.PUT("ticket", ticketRouter.updateTicket)
	r.GET("ticket/:id", ticketRouter.getTicketByID)
	r.GET("ticket/by", ticketRouter.listTicketsByUser)
	r.GET("ticket/for", ticketRouter.listTicketsForUser)
	r.DELETE("ticket/:id", ticketRouter.deleteTicket)
}

type CreateTicketRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AssignedTo  int64     `json:"assigned_to"`
	DueDate     time.Time `json:"due_date"`
}

func (tr *ticketRouter) createTicket(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	var req CreateTicketRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}
	arg := CreateTicketArgs{
		Title:       req.Title,
		Description: req.Description,
		Assigned_to: req.AssignedTo,
		Due_date:    req.DueDate,
		Created_by:  authPayload.UserID,
	}
	ticket, err := tr.ts.CreateTicket(arg, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ticket)
}

func (tr *ticketRouter) updateTicket(ctx *gin.Context) {
	var req UpdateTicketArgs
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	ticket, err := tr.ts.UpdateTicket(req, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ticket)
}

func (tr *ticketRouter) getTicketByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	ticket, err := tr.ts.GetTicket(int64(id), ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ticket)
}

func (tr *ticketRouter) listTicketsByUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	limit, offset, err := getLimitAndOffset(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	tickets, err := tr.ts.ListTicketsByUser(authPayload.UserID, offset, limit, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (tr *ticketRouter) listTicketsForUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	limit, offset, err := getLimitAndOffset(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	tickets, err := tr.ts.ListTicketsForUser(authPayload.UserID, offset, limit, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (tr *ticketRouter) deleteTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}

	/* 	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload) */
	err = tr.ts.DeleteTicket(int64(id), ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func getLimitAndOffset(ctx *gin.Context) (int32, int32, error) {
	limit, ok := ctx.GetQuery("limit")
	if !ok {
		return 0, 0, errors.New("")
	}
	offset, ok := ctx.GetQuery("offset")
	if !ok {
		return 0, 0, errors.New("")
	}
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, err
	}
	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		return 0, 0, err
	}

	return int32(limit_int), int32(offset_int), nil
}
