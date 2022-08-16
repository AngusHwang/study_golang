// Code generated by entc, DO NOT EDIT.

package partnership

import (
	"time"
)

const (
	// Label holds the string label denoting the partnership type in the database.
	Label = "partnership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the partnership in the database.
	Table = "partnerships"
)

// Columns holds all SQL columns for partnership fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCompany,
	FieldEmail,
	FieldContent,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CompanyValidator is a validator for the "company" field. It is called by the builders before save.
	CompanyValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)