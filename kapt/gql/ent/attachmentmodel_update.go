// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapt/kapt/gql/ent/attachmentmodel"
	"kapt/kapt/gql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AttachmentModelUpdate is the builder for updating AttachmentModel entities.
type AttachmentModelUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentModelMutation
}

// Where appends a list predicates to the AttachmentModelUpdate builder.
func (amu *AttachmentModelUpdate) Where(ps ...predicate.AttachmentModel) *AttachmentModelUpdate {
	amu.mutation.Where(ps...)
	return amu
}

// SetLink sets the "link" field.
func (amu *AttachmentModelUpdate) SetLink(s string) *AttachmentModelUpdate {
	amu.mutation.SetLink(s)
	return amu
}

// SetTicketID sets the "ticket_id" field.
func (amu *AttachmentModelUpdate) SetTicketID(u uuid.UUID) *AttachmentModelUpdate {
	amu.mutation.SetTicketID(u)
	return amu
}

// Mutation returns the AttachmentModelMutation object of the builder.
func (amu *AttachmentModelUpdate) Mutation() *AttachmentModelMutation {
	return amu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (amu *AttachmentModelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, amu.sqlSave, amu.mutation, amu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (amu *AttachmentModelUpdate) SaveX(ctx context.Context) int {
	affected, err := amu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (amu *AttachmentModelUpdate) Exec(ctx context.Context) error {
	_, err := amu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amu *AttachmentModelUpdate) ExecX(ctx context.Context) {
	if err := amu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (amu *AttachmentModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(attachmentmodel.Table, attachmentmodel.Columns, sqlgraph.NewFieldSpec(attachmentmodel.FieldID, field.TypeUUID))
	if ps := amu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amu.mutation.Link(); ok {
		_spec.SetField(attachmentmodel.FieldLink, field.TypeString, value)
	}
	if value, ok := amu.mutation.TicketID(); ok {
		_spec.SetField(attachmentmodel.FieldTicketID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, amu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachmentmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	amu.mutation.done = true
	return n, nil
}

// AttachmentModelUpdateOne is the builder for updating a single AttachmentModel entity.
type AttachmentModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentModelMutation
}

// SetLink sets the "link" field.
func (amuo *AttachmentModelUpdateOne) SetLink(s string) *AttachmentModelUpdateOne {
	amuo.mutation.SetLink(s)
	return amuo
}

// SetTicketID sets the "ticket_id" field.
func (amuo *AttachmentModelUpdateOne) SetTicketID(u uuid.UUID) *AttachmentModelUpdateOne {
	amuo.mutation.SetTicketID(u)
	return amuo
}

// Mutation returns the AttachmentModelMutation object of the builder.
func (amuo *AttachmentModelUpdateOne) Mutation() *AttachmentModelMutation {
	return amuo.mutation
}

// Where appends a list predicates to the AttachmentModelUpdate builder.
func (amuo *AttachmentModelUpdateOne) Where(ps ...predicate.AttachmentModel) *AttachmentModelUpdateOne {
	amuo.mutation.Where(ps...)
	return amuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (amuo *AttachmentModelUpdateOne) Select(field string, fields ...string) *AttachmentModelUpdateOne {
	amuo.fields = append([]string{field}, fields...)
	return amuo
}

// Save executes the query and returns the updated AttachmentModel entity.
func (amuo *AttachmentModelUpdateOne) Save(ctx context.Context) (*AttachmentModel, error) {
	return withHooks(ctx, amuo.sqlSave, amuo.mutation, amuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (amuo *AttachmentModelUpdateOne) SaveX(ctx context.Context) *AttachmentModel {
	node, err := amuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (amuo *AttachmentModelUpdateOne) Exec(ctx context.Context) error {
	_, err := amuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amuo *AttachmentModelUpdateOne) ExecX(ctx context.Context) {
	if err := amuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (amuo *AttachmentModelUpdateOne) sqlSave(ctx context.Context) (_node *AttachmentModel, err error) {
	_spec := sqlgraph.NewUpdateSpec(attachmentmodel.Table, attachmentmodel.Columns, sqlgraph.NewFieldSpec(attachmentmodel.FieldID, field.TypeUUID))
	id, ok := amuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AttachmentModel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := amuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachmentmodel.FieldID)
		for _, f := range fields {
			if !attachmentmodel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attachmentmodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := amuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amuo.mutation.Link(); ok {
		_spec.SetField(attachmentmodel.FieldLink, field.TypeString, value)
	}
	if value, ok := amuo.mutation.TicketID(); ok {
		_spec.SetField(attachmentmodel.FieldTicketID, field.TypeUUID, value)
	}
	_node = &AttachmentModel{config: amuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, amuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachmentmodel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	amuo.mutation.done = true
	return _node, nil
}