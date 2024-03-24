package postgres

import (
	"context"
	"log"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingEntryRepository struct {
	con *pgxpool.Pool
}

func NewPostgresWritingEntryRepository(con *pgxpool.Pool) repository.WritingEntryRepository {
	_, err := con.Exec(context.Background(), schema.WritingEntryTable)
	if err != nil {
		log.Fatalln("failed to create entry table ", err)
	}
	return &PostgresWritingEntryRepository{con: con}
}

func (r *PostgresWritingEntryRepository) Clear(ctx context.Context) error {
	_, err := r.con.Exec(context.Background(), "TRUNCATE TABLE writings;")
	return err
}

func (r *PostgresWritingEntryRepository) Add(ctx context.Context, entry model.NewWritingEntry, userId int) (int, error) {
	query := `
	INSERT INTO writings (template_id, user_id)
	VALUES ($1, $2)
	RETURNING id
	`
	var id int
	err := r.con.QueryRow(ctx, query, entry.TemplateID, userId).Scan(&id)
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
