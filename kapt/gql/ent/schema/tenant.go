package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TenantModel holds the schema definition for the TenantModel entity.
type TenantModel struct {
	ent.Schema
}

// Fields of the TenantModel.
func (TenantModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("tenant_id"),
		field.String("tenant_name").
			Unique(),
		field.Int("status"),
	}
}

// Edges of the TenantModel.
func (TenantModel) Edges() []ent.Edge {
	return nil
}
