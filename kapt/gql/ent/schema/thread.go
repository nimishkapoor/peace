package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ThreadModel holds the schema definition for the ThreadModel entity.
type ThreadModel struct {
	ent.Schema
}

// Fields of the ThreadModel.
func (ThreadModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("thread_id"),
		field.String("body").
			Unique(),
		field.Int("status"),
	}
}

// Edges of the ThreadModel.
func (ThreadModel) Edges() []ent.Edge {
	return nil
}