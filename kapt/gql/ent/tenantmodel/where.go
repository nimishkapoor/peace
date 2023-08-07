// Code generated by ent, DO NOT EDIT.

package tenantmodel

import (
	"kapt/kapt/gql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLTE(FieldID, id))
}

// TenantName applies equality check predicate on the "tenant_name" field. It's identical to TenantNameEQ.
func TenantName(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldTenantName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldStatus, v))
}

// TenantNameEQ applies the EQ predicate on the "tenant_name" field.
func TenantNameEQ(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldTenantName, v))
}

// TenantNameNEQ applies the NEQ predicate on the "tenant_name" field.
func TenantNameNEQ(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNEQ(FieldTenantName, v))
}

// TenantNameIn applies the In predicate on the "tenant_name" field.
func TenantNameIn(vs ...string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldIn(FieldTenantName, vs...))
}

// TenantNameNotIn applies the NotIn predicate on the "tenant_name" field.
func TenantNameNotIn(vs ...string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNotIn(FieldTenantName, vs...))
}

// TenantNameGT applies the GT predicate on the "tenant_name" field.
func TenantNameGT(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGT(FieldTenantName, v))
}

// TenantNameGTE applies the GTE predicate on the "tenant_name" field.
func TenantNameGTE(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGTE(FieldTenantName, v))
}

// TenantNameLT applies the LT predicate on the "tenant_name" field.
func TenantNameLT(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLT(FieldTenantName, v))
}

// TenantNameLTE applies the LTE predicate on the "tenant_name" field.
func TenantNameLTE(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLTE(FieldTenantName, v))
}

// TenantNameContains applies the Contains predicate on the "tenant_name" field.
func TenantNameContains(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldContains(FieldTenantName, v))
}

// TenantNameHasPrefix applies the HasPrefix predicate on the "tenant_name" field.
func TenantNameHasPrefix(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldHasPrefix(FieldTenantName, v))
}

// TenantNameHasSuffix applies the HasSuffix predicate on the "tenant_name" field.
func TenantNameHasSuffix(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldHasSuffix(FieldTenantName, v))
}

// TenantNameEqualFold applies the EqualFold predicate on the "tenant_name" field.
func TenantNameEqualFold(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEqualFold(FieldTenantName, v))
}

// TenantNameContainsFold applies the ContainsFold predicate on the "tenant_name" field.
func TenantNameContainsFold(v string) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldContainsFold(FieldTenantName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.TenantModel {
	return predicate.TenantModel(sql.FieldLTE(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TenantModel) predicate.TenantModel {
	return predicate.TenantModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TenantModel) predicate.TenantModel {
	return predicate.TenantModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TenantModel) predicate.TenantModel {
	return predicate.TenantModel(func(s *sql.Selector) {
		p(s.Not())
	})
}
