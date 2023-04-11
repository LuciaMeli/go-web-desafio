package tickets

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket,error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
    return s.repo.GetAll(ctx)
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
    t, err := s.repo.GetTicketByDestination(ctx, destination)
	if err!= nil {
        return []domain.Ticket{}, err
    }
	return t, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	list, err := s.repo.GetAll(ctx)
	if err!= nil {
        return 0, err
    }
	counter := 0
    for _, t := range list {
        if t.Country == destination {
            counter++
        }
    }
    return (float64(counter) / float64(len(list))), nil
}

