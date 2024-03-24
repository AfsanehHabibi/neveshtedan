package model

type WritingTemplate struct {
	Id          int
	UserId      int
	Title       string
	Description string
}

type WritingTemplateField struct {
	Name       string
	TemplateId int
	Type       FieldType
}

type FieldType string

const (
	TEXT   FieldType = "TEXT"
	NUMBER FieldType = "NUMBER"
	IMAGE  FieldType = "IMAGE"
	VIDEO  FieldType = "VIDEO"
)
