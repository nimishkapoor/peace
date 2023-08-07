package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// UserModel holds the schema definition for the UserModel entity.
type UserModel struct {
	ent.Schema
}

// Fields of the UserModel.
func (UserModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("user_id"),
		field.String("user_name").
			Unique(),
		field.String("pswd"),
		field.UUID("tenant_id", uuid.UUID{}),
		field.String("role"),
	}
}

// Edges of the UserModel.
func (UserModel) Edges() []ent.Edge {
	return nil
}
