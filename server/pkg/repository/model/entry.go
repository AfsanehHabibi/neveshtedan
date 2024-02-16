package model

type WritingEntry struct {
	Id int
	TemplateId int
	UserId int
}

type WritingEntryField struct {
	Name string
	Value *string
	EntryId int
}
