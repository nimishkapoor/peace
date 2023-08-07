// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kapt/kapt/gql/ent/categorymodel"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CategoryModel is the model entity for the CategoryModel schema.
type CategoryModel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID     uuid.UUID `json:"tenant_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CategoryModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case categorymodel.FieldName:
			values[i] = new(sql.NullString)
		case categorymodel.FieldID, categorymodel.FieldTenantID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CategoryModel fields.
func (cm *CategoryModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case categorymodel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cm.ID = *value
			}
		case categorymodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cm.Name = value.String
			}
		case categorymodel.FieldTenantID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value != nil {
				cm.TenantID = *value
			}
		default:
			cm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CategoryModel.
// This includes values selected through modifiers, order, etc.
func (cm *CategoryModel) Value(name string) (ent.Value, error) {
	return cm.selectValues.Get(name)
}

// Update returns a builder for updating this CategoryModel.
// Note that you need to call CategoryModel.Unwrap() before calling this method if this CategoryModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (cm *CategoryModel) Update() *CategoryModelUpdateOne {
	return NewCategoryModelClient(cm.config).UpdateOne(cm)
}

// Unwrap unwraps the CategoryModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cm *CategoryModel) Unwrap() *CategoryModel {
	_tx, ok := cm.config.driver.(*txDriver)
	if !ok {
		panic("ent: CategoryModel is not a transactional entity")
	}
	cm.config.driver = _tx.drv
	return cm
}

// String implements the fmt.Stringer.
func (cm *CategoryModel) String() string {
	var builder strings.Builder
	builder.WriteString("CategoryModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cm.ID))
	builder.WriteString("name=")
	builder.WriteString(cm.Name)
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", cm.TenantID))
	builder.WriteByte(')')
	return builder.String()
}

// CategoryModels is a parsable slice of CategoryModel.
type CategoryModels []*CategoryModel