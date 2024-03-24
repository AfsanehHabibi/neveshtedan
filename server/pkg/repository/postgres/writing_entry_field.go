package postgres

import (
	"context"
	"log"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/schema"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresWritingEntryFieldRepository struct {
	con *pgxpool.Pool
}

func NewPostgresWritingEntryFieldRepository(con *pgxpool.Pool) repository.WritingEntryFieldRepository {
	_, err := con.Exec(context.Background(), schema.WritingEntryImageFieldTable)
	if err != nil {
		log.Fatalln("failed to create entry image field table ", err)
	}
	_, err = con.Exec(context.Background(), schema.WritingEntryNumberFieldTable)
	if err != nil {
		log.Fatalln("failed to create entry number field table ", err)
	}
	_, err = con.Exec(context.Background(), schema.WritingEntryTextFieldTable)
	if err != nil {
		log.Fatalln("failed to create entry text field table ", err)
	}
	_, err = con.Exec(context.Background(), schema.WritingEntryVideoFieldTable)
	if err != nil {
		log.Fatalln("failed to create entry video field table ", err)
	}
	return &PostgresWritingEntryFieldRepository{con: con}
}

func (r *PostgresWritingEntryFieldRepository) Add(ctx context.Context, entryId int, field model.NewWritingEntryField) error {
	switch field.Type {
	case model.FieldTypeImage:
		if field.URL == nil {
			break
		}
		query := `
			INSERT INTO writing_image_fields (name, url, entry_id)
			VALUES ($1, $2, $3)
		`
		_, err := r.con.Exec(ctx, query, field.Name, *field.URL, entryId)
		if err != nil {
			return err
		}
		return nil
	case model.FieldTypeNumber:
		if field.Number == nil {
			break
		}
		query := `
			INSERT INTO writing_number_fields (name, value, entry_id)
			VALUES ($1, $2, $3)
		`
		_, err := r.con.Exec(ctx, query, field.Name, *field.Number, entryId)
		if err != nil {
			return err
		}
		return nil
	case model.FieldTypeText:
		if field.Text == nil {
			break
		}
		query := `
			INSERT INTO writing_text_fields (name, text, entry_id)
			VALUES ($1, $2, $3)
		`
		_, err := r.con.Exec(ctx, query, field.Name, *field.Text, entryId)
		if err != nil {
			return err
		}
		return nil
	case model.FieldTypeVideo:
		if field.URL == nil {
			break
		}
		query := `
			INSERT INTO writing_video_fields (name, url, entry_id)
			VALUES ($1, $2, $3)
		`
		_, err := r.con.Exec(ctx, query, field.Name, *field.URL, entryId)
		if err != nil {
			return err
		}
		return nil
	}
	//it should not reach here
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
	var result []model.WritingEntryField
	textFields, err := r.getAllTexts(ctx, id)
	if err != nil {
		return result, err
	}
	result = append(result, textFields...)
	imageFields, err := r.getAllImages(ctx, id)
	if err != nil {
		return result, err
	}
	result = append(result, imageFields...)
	numberFields, err := r.getAllNumbers(ctx, id)
	if err != nil {
		return result, err
	}
	result = append(result, numberFields...)
	videoFields, err := r.getAllVideos(ctx, id)
	if err != nil {
		log.Println("video len ", len(videoFields))
		return result, err
	}
	result = append(result, videoFields...)
	log.Println("video len ", len(videoFields))
	return result, nil
}

func (r *PostgresWritingEntryFieldRepository) getAllTexts(ctx context.Context, id int) ([]model.WritingEntryField, error) {
	query := `
	SELECT name, text
	FROM writing_text_fields
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
		var text string
		err := rows.Scan(&entry.Name, &text)
		entry.Value = model.TextValue{Text: text}
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

func (r *PostgresWritingEntryFieldRepository) getAllImages(ctx context.Context, id int) ([]model.WritingEntryField, error) {
	query := `
	SELECT name, url
	FROM writing_image_fields
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
		var url string
		err := rows.Scan(&entry.Name, &url)
		entry.Value = model.ImageValue{URL: url}
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

func (r *PostgresWritingEntryFieldRepository) getAllVideos(ctx context.Context, id int) ([]model.WritingEntryField, error) {
	query := `
	SELECT name, url
	FROM writing_video_fields
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
		var url string
		err := rows.Scan(&entry.Name, &url)
		entry.Value = model.VideoValue{URL: url}
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

func (r *PostgresWritingEntryFieldRepository) getAllNumbers(ctx context.Context, id int) ([]model.WritingEntryField, error) {
	query := `
	SELECT name, value
	FROM writing_number_fields
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
		var number float64
		err := rows.Scan(&entry.Name, &number)
		entry.Value = model.NumberValue{Number: number}
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
	_, err := r.con.Exec(context.Background(), "TRUNCATE TABLE writing_image_fields;")
	if err != nil {
		return err
	}
	_, err = r.con.Exec(context.Background(), "TRUNCATE TABLE writing_text_fields;")
	if err != nil {
		return err
	}
	_, err = r.con.Exec(context.Background(), "TRUNCATE TABLE writing_number_fields;")
	if err != nil {
		return err
	}
	_, err = r.con.Exec(context.Background(), "TRUNCATE TABLE writing_video_fields;")
	return err
}
