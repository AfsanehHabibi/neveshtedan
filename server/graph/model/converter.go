package model

func ConvertNewToWritingEntryField(input NewWritingEntryField) *WritingEntryField {
    return &WritingEntryField{
        Name:  input.Name,
        Value: TextValue{*input.Text},
    }
}
