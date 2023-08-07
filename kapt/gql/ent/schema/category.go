package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CategoryModel holds the schema definition for the CategoryModel entity.
type CategoryModel struct {
	ent.Schema
}

// Fields of the CategoryModel.
func (CategoryModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("category_id"),
		field.String("name"),
		field.UUID("tenant_id", uuid.UUID{}),
	}
}

// Edges of the CategoryModel.
func (CategoryModel) Edges() []ent.Edge {
	return nil
}
