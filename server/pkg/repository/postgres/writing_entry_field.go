package postgres

import (
	"context"
	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingEntryFieldRepository struct {
	con *pgxpool.Pool
}

func NewPostgresWritingEntryFieldRepository(con *pgxpool.Pool) repository.WritingEntryFieldRepository {
	return &PostgresWritingEntryFieldRepository{con: con}
}

func (r *PostgresWritingEntryFieldRepository) Add(ctx context.Context, entryId int, field model.NewWritingEntryField) error {
	if field.Value == nil {
		return nil
	}

	value := field.Value
	query := `
	INSERT INTO writing_fields (name, value, entry_id)
	VALUES ($1, $2, $3)
	`
	_, err := r.con.Exec(ctx, query, field.Name, value, entryId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresWritingEntryFieldRepository) AddAll(ctx context.Context, entryId int, fields []model.NewWritingEntryField) error {
	//TODO(use bulk insert)
	for _, field := range fields {
		err := r.Add(ctx, entryId, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresWritingEntryFieldRepository) GetAll(ctx context.Context, id int) ([]model.WritingEntryField, error) {
	query := `
	SELECT name, value
	FROM writing_fields
	WHERE entry_id = $1
	`
	rows, err := r.con.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.WritingEntryField

	for rows.Next() {
		var entry model.WritingEntryField
		err := rows.Scan(&entry.Name, &entry.Value)
		if err != nil {
			return nil, err
		}
		result = append(result, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PostgresWritingEntryFieldRepository) Clear(ctx context.Context) error {
	_, err := r.con.Exec(context.Background(), "TRUNCATE TABLE writing_fields;")
	return err
}
