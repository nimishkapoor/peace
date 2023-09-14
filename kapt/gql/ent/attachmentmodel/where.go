// Code generated by ent, DO NOT EDIT.

package attachmentmodel

import (
	"kapt/kapt/gql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLTE(FieldID, id))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldLink, v))
}

// ThreadID applies equality check predicate on the "thread_id" field. It's identical to ThreadIDEQ.
func ThreadID(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldThreadID, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldContainsFold(FieldLink, v))
}

// ThreadIDEQ applies the EQ predicate on the "thread_id" field.
func ThreadIDEQ(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldEQ(FieldThreadID, v))
}

// ThreadIDNEQ applies the NEQ predicate on the "thread_id" field.
func ThreadIDNEQ(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNEQ(FieldThreadID, v))
}

// ThreadIDIn applies the In predicate on the "thread_id" field.
func ThreadIDIn(vs ...uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldIn(FieldThreadID, vs...))
}

// ThreadIDNotIn applies the NotIn predicate on the "thread_id" field.
func ThreadIDNotIn(vs ...uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldNotIn(FieldThreadID, vs...))
}

// ThreadIDGT applies the GT predicate on the "thread_id" field.
func ThreadIDGT(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGT(FieldThreadID, v))
}

// ThreadIDGTE applies the GTE predicate on the "thread_id" field.
func ThreadIDGTE(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldGTE(FieldThreadID, v))
}

// ThreadIDLT applies the LT predicate on the "thread_id" field.
func ThreadIDLT(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLT(FieldThreadID, v))
}

// ThreadIDLTE applies the LTE predicate on the "thread_id" field.
func ThreadIDLTE(v uuid.UUID) predicate.AttachmentModel {
	return predicate.AttachmentModel(sql.FieldLTE(FieldThreadID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AttachmentModel) predicate.AttachmentModel {
	return predicate.AttachmentModel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AttachmentModel) predicate.AttachmentModel {
	return predicate.AttachmentModel(func(s *sql.Selector) {
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
func Not(p predicate.AttachmentModel) predicate.AttachmentModel {
	return predicate.AttachmentModel(func(s *sql.Selector) {
		p(s.Not())
	})
}
