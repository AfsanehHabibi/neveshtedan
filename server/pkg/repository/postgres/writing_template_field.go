package postgres

import (
	"context"
	"log"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingTemplateFieldRepository struct {
	con *pgxpool.Pool
}

func NewPostgresWritingTemplateFieldRepository(con *pgxpool.Pool) repository.WritingTemplateFieldRepository {
    _, err := con.Exec(context.Background(), schema.WritingTemplateFieldTable)
	if err != nil {
		log.Fatalln("failed to create template field table ", err)
	}
	return &PostgresWritingTemplateFieldRepository{con: con}
}

func (r *PostgresWritingTemplateFieldRepository) AddAll(ctx context.Context, templateID int, fields []model.NewWritingTemplateField) error {
    tx, err := r.con.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    query := `
        INSERT INTO template_fields (template_id, name, description, type)
        VALUES ($1, $2, $3, $4)
    `

    for _, field := range fields {
        _, err := tx.Exec(ctx, query, templateID, field.Name, field.Description, field.Type)
        if err != nil {
            return err
        }
    }

    if err := tx.Commit(ctx); err != nil {
        return err
    }

    return nil
}

func (r *PostgresWritingTemplateFieldRepository) GetAll(ctx context.Context, templateID int) ([]model.WritingTemplateField, error) {
    query := `
        SELECT name, description, type
		FROM template_fields
        WHERE template_id = $1
    `

    rows, err := r.con.Query(ctx, query, templateID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var fields []model.WritingTemplateField
    for rows.Next() {
        var field model.WritingTemplateField
        if err := rows.Scan(&field.Name, &field.Description, &field.Type); err != nil {
            return nil, err
        }
        fields = append(fields, field)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return fields, nil
}

func (r *PostgresWritingTemplateFieldRepository) Clear(ctx context.Context) error {
	return util.ClearTable(ctx, "template_fields", r.con)
}
