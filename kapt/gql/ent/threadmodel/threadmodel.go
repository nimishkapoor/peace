// Code generated by ent, DO NOT EDIT.

package threadmodel

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the threadmodel type in the database.
	Label = "thread_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "thread_id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the threadmodel in the database.
	Table = "thread_models"
)

// Columns holds all SQL columns for threadmodel fields.
var Columns = []string{
	FieldID,
	FieldBody,
	FieldStatus,
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

// OrderOption defines the ordering options for the ThreadModel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}