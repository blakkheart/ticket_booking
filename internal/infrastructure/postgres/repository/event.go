package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	eventModel "ticket-booking/internal/event"

	"github.com/jackc/pgx/v5/pgxpool"
)

type eventRepository struct {
	queries *sqlc_repository.Queries
	logger  *slog.Logger
}

func NewEventRepository(db *pgxpool.Pool, logger *slog.Logger) *eventRepository {
	return &eventRepository{
		queries: sqlc_repository.New(db),
		logger:  logger,
	}
}

func (repo *eventRepository) Create(ctx context.Context, e *eventModel.EventIn) (*eventModel.Event, error) {
	event, err := repo.queries.CreateEvent(
		ctx,
		sqlc_repository.CreateEventParams{
			Title:       e.Title,
			Description: e.Description,
			Location:    e.Location,
			StartsAt:    e.StartsAt,
			EndsAt:      e.EndsAt,
			Status:      e.Status,
		},
	)

	return repo.fromSqlcEvent(&event), err
}

func (repo *eventRepository) Delete(id int64) error {
	return nil
}

func (repo *eventRepository) Get(id int64) (*eventModel.Event, error) {
	return nil, nil
}

func (repo *eventRepository) GetMany(filter any) ([]*eventModel.Event, error) {
	return nil, nil
}

func (repo *eventRepository) fromSqlcEvent(e *sqlc_repository.Event) *eventModel.Event {
	event := &eventModel.Event{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Location:    e.Location,
		StartsAt:    e.StartsAt,
		EndsAt:      e.EndsAt,
		Status:      e.Status,
	}
	return event
}
