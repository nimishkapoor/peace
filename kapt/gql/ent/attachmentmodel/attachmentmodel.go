// Code generated by ent, DO NOT EDIT.

package attachmentmodel

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the attachmentmodel type in the database.
	Label = "attachment_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "attachment_id"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldThreadID holds the string denoting the thread_id field in the database.
	FieldThreadID = "thread_id"
	// Table holds the table name of the attachmentmodel in the database.
	Table = "attachment_models"
)

// Columns holds all SQL columns for attachmentmodel fields.
var Columns = []string{
	FieldID,
	FieldLink,
	FieldThreadID,
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

// OrderOption defines the ordering options for the AttachmentModel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}

// ByThreadID orders the results by the thread_id field.
func ByThreadID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreadID, opts...).ToFunc()
}
