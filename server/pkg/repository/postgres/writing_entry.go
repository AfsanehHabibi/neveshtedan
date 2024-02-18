package postgres

import (
	"context"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingEntryRepository struct {
	con *pgxpool.Pool
}

// Clear implements repository.WritingEntryRepository.
func (*PostgresWritingEntryRepository) Clear(ctx context.Context) error {
	panic("unimplemented")
}

// Add implements repository.WritingEntryRepository.
func (r *PostgresWritingEntryRepository) Add(ctx context.Context, entry model.NewWritingEntry) (int, error) {
	query := `
	INSERT INTO writings (template_id, user_id)
	VALUES ($1, $2)
	RETURNING id
	`
	var id int
	err := r.con.QueryRow(ctx, query, entry.TemplateID, entry.UserID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresWritingEntryRepository) GetAll(ctx context.Context) ([]model.WritingEntry, error) {
	query := `
	SELECT id, template_id, user_id
	FROM writings
	`
	rows, err := r.con.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.WritingEntry
	for rows.Next() {
		var entry model.WritingEntry
		if err := rows.Scan(&entry.ID, &entry.TemplateID, &entry.UserID); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *PostgresWritingEntryRepository) GetById(ctx context.Context, id int) (*model.WritingEntry, error) {
	query := `
	SELECT id, template_id, user_id
	FROM writings
	WHERE id = $1
	`
	var entry model.WritingEntry
	err := r.con.QueryRow(ctx, query, id).Scan(&entry.ID, &entry.TemplateID, &entry.UserID)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &entry, nil
}

func NewPostgresWritingEntryRepository() repository.WritingEntryRepository {
	return &PostgresWritingEntryRepository{}
}
