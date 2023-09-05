package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttachmentModel holds the schema definition for the AttachmentModel entity.
type AttachmentModel struct {
	ent.Schema
}

// Fields of the AttachmentModel.
func (AttachmentModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("attachment_id"),
		field.String("link"),
		field.UUID("thread_id", uuid.UUID{}),
	}
}

// Edges of the AttachmentModel.
func (AttachmentModel) Edges() []ent.Edge {
	return nil
}
