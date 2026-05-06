package repository

import (
	"context"
	"log/slog"
	sqlc_repository "ticket-booking/internal/db/sqlc"
	"ticket-booking/internal/event/models"

	"github.com/google/uuid"
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

func NewEventRepositoryFromQueries(q *sqlc_repository.Queries, logger *slog.Logger) *eventRepository {
	return &eventRepository{
		queries: q,
		logger:  logger,
	}
}

func (repo *eventRepository) Create(ctx context.Context, e *models.EventIn) (*models.Event, error) {
	event, err := repo.queries.CreateEvent(
		ctx,
		sqlc_repository.CreateEventParams{
			ID:          e.ID,
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

func (repo *eventRepository) Delete(id uuid.UUID) error {
	return nil
}

func (repo *eventRepository) Get(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	event, err := repo.queries.GetEvent(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return repo.fromSqlcEvent(&event), nil
}

func (repo *eventRepository) GetMany(ctx context.Context, filter *models.EventFilter) ([]*models.Event, error) {
	return nil, nil
}

func (repo *eventRepository) UpdateById(
	ctx context.Context,
	id uuid.UUID,
	e *models.EventUpdate,
) (*models.Event, error) {
	return nil, nil
}

func (repo *eventRepository) fromSqlcEvent(e *sqlc_repository.Event) *models.Event {
	event := &models.Event{
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
