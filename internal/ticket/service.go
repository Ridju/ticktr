package ticket

import (
	"context"

	db "github.com/Ridju/ticktr/internal/db/sqlc"
)

type ITicketService interface {
	CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error)
	UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error)
	GetTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error)
	DeleteTicket(ID int64, ctx context.Context) error
}

type TicketService struct {
	repo ITicketRepository
}

func NewTicketService(repo ITicketRepository) ITicketRepository {
	return &TicketRepository{
		repo: repo,
	}
}

func (s *TicketService) CreateTicket(args CreateTicketArgs, ctx context.Context) (db.Ticket, error) {
	ticket, err := s.repo.CreateTicket(args, ctx)
	if err != nil {
		return db.Ticket{}, nil
	}

	return ticket, err
}

func (s *TicketService) UpdateTicket(args UpdateTicketArgs, ctx context.Context) (db.Ticket, error) {
	ticket, err := s.repo.UpdateTicket(args, ctx)
	if err != nil {
		return db.Ticket{}, nil
	}
	return ticket, err
}

func (s *TicketService) GetTickets(offset int32, page_size int32, ctx context.Context) ([]db.Ticket, error) {
	tickets, err := s.repo.GetTickets(offset, page_size, ctx)
	if err != nil {
		return []db.Ticket{}, nil
	}
	return tickets, err
}

func (s *TicketService) DeleteTicket(ID int64, ctx context.Context) error {
	if err := s.repo.DeleteTicket(ID, ctx); err != nil {
		return err
	}
	return nil
}
