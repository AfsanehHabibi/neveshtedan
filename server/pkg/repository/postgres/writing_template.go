package postgres

import (
	"context"
	"log"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingTemplateRepository struct {
	con *pgxpool.Pool
}

func NewPostgresWritingTemplateRepository(con *pgxpool.Pool) repository.WritingTemplateRepository {
	_, err := con.Exec(context.Background(), schema.WritingTemplateTable)
	if err != nil {
		log.Fatalln("failed to create template table ", err)
	}
	return &PostgresWritingTemplateRepository{con: con}
}

func (r *PostgresWritingTemplateRepository) Add(ctx context.Context, userId int, template model.NewWritingTemplate) (int, error) {
	query := `
	INSERT INTO templates (user_id, title, description)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var id int
	err := r.con.QueryRow(ctx, query, userId, template.Title, template.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresWritingTemplateRepository) GetAll(ctx context.Context) ([]model.WritingTemplate, error) {
	query := `
	SELECT id, user_id, title, description
	FROM templates
	`
	rows, err := r.con.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []model.WritingTemplate
	for rows.Next() {
		var template model.WritingTemplate
		if err := rows.Scan(&template.ID, &template.UserID, &template.Title, &template.Description); err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return templates, nil
}

func (r *PostgresWritingTemplateRepository) GetById(ctx context.Context, id int) (*model.WritingTemplate, error) {
	query := `
	SELECT id, user_id, title, description
	FROM templates
	WHERE id = $1
	`
	var template model.WritingTemplate
	err := r.con.QueryRow(ctx, query, id).Scan(&template.ID, &template.UserID, &template.Title, &template.Description)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &template, nil
}

func (r *PostgresWritingTemplateRepository) Clear(ctx context.Context) error {
	return util.ClearTable(ctx, "templates", r.con)
}
