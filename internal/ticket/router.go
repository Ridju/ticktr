package ticket

import (
	"github.com/Ridju/ticktr/config"
	db "github.com/Ridju/ticktr/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ticketRouter struct {
	repo   ITicketRepository
	ts     ITicketService
	config config.Config
}

func InitTicketRouter(r *gin.RouterGroup, store db.Store, config config.Config) {
	ticketRepo := NewTicketRepository(store)
	ticketService := NewTicketService(ticketRepo)

	ticketRouter := ticketRouter{
		repo:   ticketRepo,
		ts:     ticketService,
		config: config,
	}

	r.POST("", ticketRouter.createTicket)
	r.PUT("", ticketRouter.updateTicket)
	r.GET("/:id", ticketRouter.getTicket)
	r.GET("", ticketRouter.getTickets)
	r.DELETE("/:id", ticketRouter.deleteTicket)
}

func (tr *ticketRouter) createTicket(ctx *gin.Context) {

}
func (tr *ticketRouter) updateTicket(ctx *gin.Context) {}
func (tr *ticketRouter) getTicket(ctx *gin.Context)    {}
func (tr *ticketRouter) getTickets(ctx *gin.Context)   {}
func (tr *ticketRouter) deleteTicket(ctx *gin.Context) {}
