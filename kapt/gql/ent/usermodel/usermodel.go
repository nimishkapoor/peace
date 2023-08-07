// Code generated by ent, DO NOT EDIT.

package usermodel

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the usermodel type in the database.
	Label = "user_model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "user_id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldPswd holds the string denoting the pswd field in the database.
	FieldPswd = "pswd"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// Table holds the table name of the usermodel in the database.
	Table = "user_models"
)

// Columns holds all SQL columns for usermodel fields.
var Columns = []string{
	FieldID,
	FieldUserName,
	FieldPswd,
	FieldTenantID,
	FieldRole,
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

// OrderOption defines the ordering options for the UserModel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserName orders the results by the user_name field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByPswd orders the results by the pswd field.
func ByPswd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPswd, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}
