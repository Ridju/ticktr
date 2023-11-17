package ticket

import (
	"context"
	"time"

	db "github.com/Ridju/ticktr/internal/db/sqlc"
)

type ITicketRepository interface {
	CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error)
	UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error)
	GetTicket(ID int64, ctx context.Context) (db.Ticket, error)
	ListTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error)
	ListTicketsForUser(UserID int64, offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error)
	ListTicketsByUser(UserID int64, offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error)
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
	Assigned_to int64
	Created_by  int64
}

func (r *TicketRepository) GetTicket(ID int64, ctx context.Context) (db.Ticket, error) {
	ticket, err := r.store.GetTicketByID(ctx, ID)
	if err != nil {
		return db.Ticket{}, nil
	}
	return ticket, err
}

func (r *TicketRepository) CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error) {
	arg := db.CreateTicketParams{
		Title:       args.Title,
		Description: args.Description,
		AssignedTo:  args.Assigned_to,
		CreatedBy:   args.Created_by,
		DueDate:     args.Due_date,
	}

	ticket, err := r.store.CreateTicket(ctx, arg)
	if err != nil {
		return db.Ticket{}, nil
	}

	return ticket, err
}

type UpdateTicketArgs struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Assigned_to int64     `json:"assigned_to"`
	CreatedBy   int64     `json:"created_by"`
	DueDate     time.Time `json:"due_date"`
}

func (r *TicketRepository) UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error) {
	arg := db.UpdateTicketParams{
		ID:          args.ID,
		Title:       args.Title,
		Description: args.Description,
		AssignedTo:  args.ID,
		CreatedBy:   args.ID,
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

func (r *TicketRepository) ListTicketsForUser(UserID int64, offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error) {
	arg := db.ListTicketsForUserParams{
		AssignedTo: UserID,
		Limit:      page_size,
		Offset:     offset,
	}
	tickets, err := r.store.ListTicketsForUser(ctx, arg)
	if err != nil {
		return []db.Ticket{}, err
	}

	return tickets, err
}

func (r *TicketRepository) ListTicketsByUser(UserID int64, offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error) {
	arg := db.ListTicketsByUserParams{
		CreatedBy: UserID,
		Limit:     page_size,
		Offset:    offset,
	}
	tickets, err := r.store.ListTicketsByUser(ctx, arg)
	if err != nil {
		return []db.Ticket{}, err
	}

	return tickets, err
}

func (r *TicketRepository) ListTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error) {
	arg := db.ListTicketsParams{
		Limit:  page_size,
		Offset: offset,
	}
	tickets, err := r.store.ListTickets(ctx, arg)
	if err != nil {
		return []db.Ticket{}, nil
	}

	return tickets, err
}
