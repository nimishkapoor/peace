package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketModel holds the schema definition for the TicketModel entity.
type TicketModel struct {
	ent.Schema
}

// Fields of the TicketModel.
func (TicketModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("ticket_id"),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("tenant_id", uuid.UUID{}),
		field.String("subject"),
		field.String("body"),
		field.UUID("category", uuid.UUID{}),
		field.String("label"),
		field.UUID("assignee_id", uuid.UUID{}),
		field.Int("severity"),
		field.Int("status"),
		field.Time("time"),
	}
}

// Edges of the TicketModel.
func (TicketModel) Edges() []ent.Edge {
	return nil
}
