// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kapt/kapt/gql/ent/threadmodel"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ThreadModel is the model entity for the ThreadModel schema.
type ThreadModel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Time holds the value of the "time" field.
	Time time.Time `json:"time,omitempty"`
	// TicketUUID holds the value of the "ticket_uuid" field.
	TicketUUID uuid.UUID `json:"ticket_uuid,omitempty"`
	// Source holds the value of the "source" field.
	Source       string `json:"source,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThreadModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case threadmodel.FieldBody, threadmodel.FieldSource:
			values[i] = new(sql.NullString)
		case threadmodel.FieldTime:
			values[i] = new(sql.NullTime)
		case threadmodel.FieldID, threadmodel.FieldTicketUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThreadModel fields.
func (tm *ThreadModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case threadmodel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tm.ID = *value
			}
		case threadmodel.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				tm.Body = value.String
			}
		case threadmodel.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				tm.Time = value.Time
			}
		case threadmodel.FieldTicketUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_uuid", values[i])
			} else if value != nil {
				tm.TicketUUID = *value
			}
		case threadmodel.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				tm.Source = value.String
			}
		default:
			tm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ThreadModel.
// This includes values selected through modifiers, order, etc.
func (tm *ThreadModel) Value(name string) (ent.Value, error) {
	return tm.selectValues.Get(name)
}

// Update returns a builder for updating this ThreadModel.
// Note that you need to call ThreadModel.Unwrap() before calling this method if this ThreadModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (tm *ThreadModel) Update() *ThreadModelUpdateOne {
	return NewThreadModelClient(tm.config).UpdateOne(tm)
}

// Unwrap unwraps the ThreadModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tm *ThreadModel) Unwrap() *ThreadModel {
	_tx, ok := tm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThreadModel is not a transactional entity")
	}
	tm.config.driver = _tx.drv
	return tm
}

// String implements the fmt.Stringer.
func (tm *ThreadModel) String() string {
	var builder strings.Builder
	builder.WriteString("ThreadModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tm.ID))
	builder.WriteString("body=")
	builder.WriteString(tm.Body)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(tm.Time.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ticket_uuid=")
	builder.WriteString(fmt.Sprintf("%v", tm.TicketUUID))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(tm.Source)
	builder.WriteByte(')')
	return builder.String()
}

// ThreadModels is a parsable slice of ThreadModel.
type ThreadModels []*ThreadModel
