package ticket

import (
	"context"
	"time"

	db "github.com/Ridju/ticktr/internal/db/sqlc"
)

type ITicketRepository interface {
	CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error)
	UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error)
	GetTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error)
	DeleteTicket(ID int64, ctx context.Context) error
}

type TicketRepository struct {
	store db.Store
}

func NewTicketRepository(store db.Store) ITicketRepository {
	return &TicketRepository{
		store: store,
	}
}

type CreateTicketArgs struct {
	Title       string
	Description string
	Due_date    time.Time
	Assigned_to db.User
	Created_by  db.User
}

func (r *TicketRepository) CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error) {
	arg := db.CreateTicketParams{
		Title:       args.Title,
		Description: args.Description,
		AssignedTo:  args.Assigned_to.ID,
		CreatedBy:   args.Created_by.ID,
		DueDate:     args.Due_date,
	}

	ticket, err := r.store.CreateTicket(ctx, arg)
	if err != nil {
		return db.Ticket{}, nil
	}

	return ticket, err
}

type UpdateTicketArgs struct {
	ID          int64
	Title       string
	Description string
	Assigned_to db.User
	CreatedBy   db.User
	DueDate     time.Time
}

func (r *TicketRepository) UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error) {
	arg := db.UpdateTicketParams{
		ID:          args.ID,
		Title:       args.Title,
		Description: args.Description,
		AssignedTo:  args.Assigned_to.ID,
		CreatedBy:   args.CreatedBy.ID,
		DueDate:     args.DueDate,
	}

	ticket, err := r.store.UpdateTicket(ctx, arg)
	if err != nil {
		return db.Ticket{}, nil
	}
	return ticket, err
}

func (r *TicketRepository) GetTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error) {
	arg := db.ListTicketsParams{
		Limit:  page_size,
		Offset: (offset - 1) * page_size,
	}
	tickets, err := r.store.ListTickets(ctx, arg)
	if err != nil {
		return []db.Ticket{}, nil
	}

	return tickets, err
}

func (r *TicketRepository) DeleteTicket(ID int64, ctx context.Context) error {
	if err := r.store.DeleteTicket(ctx, ID); err != nil {
		return err
	}
	return nil
}
