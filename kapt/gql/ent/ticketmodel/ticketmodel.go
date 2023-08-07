// Code generated by ent, DO NOT EDIT.

package ticketmodel

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ticketmodel type in the database.
	Label = "ticket_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ticket_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldLabel holds the string denoting the label field in the database.
	FieldLabel = "label"
	// FieldAssigneeID holds the string denoting the assignee_id field in the database.
	FieldAssigneeID = "assignee_id"
	// FieldSeverity holds the string denoting the severity field in the database.
	FieldSeverity = "severity"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// Table holds the table name of the ticketmodel in the database.
	Table = "ticket_models"
)

// Columns holds all SQL columns for ticketmodel fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTenantID,
	FieldSubject,
	FieldBody,
	FieldCategory,
	FieldLabel,
	FieldAssigneeID,
	FieldSeverity,
	FieldStatus,
	FieldTime,
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

// OrderOption defines the ordering options for the TicketModel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByLabel orders the results by the label field.
func ByLabel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLabel, opts...).ToFunc()
}

// ByAssigneeID orders the results by the assignee_id field.
func ByAssigneeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAssigneeID, opts...).ToFunc()
}

// BySeverity orders the results by the severity field.
func BySeverity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSeverity, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}
