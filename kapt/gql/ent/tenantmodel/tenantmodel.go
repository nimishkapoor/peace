// Code generated by ent, DO NOT EDIT.

package tenantmodel

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tenantmodel type in the database.
	Label = "tenant_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "tenant_id"
	// FieldTenantName holds the string denoting the tenant_name field in the database.
	FieldTenantName = "tenant_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the tenantmodel in the database.
	Table = "tenant_models"
)

// Columns holds all SQL columns for tenantmodel fields.
var Columns = []string{
	FieldID,
	FieldTenantName,
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

// OrderOption defines the ordering options for the TenantModel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTenantName orders the results by the tenant_name field.
func ByTenantName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
