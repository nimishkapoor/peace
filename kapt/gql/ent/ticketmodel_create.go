// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapt/kapt/gql/ent/ticketmodel"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketModelCreate is the builder for creating a TicketModel entity.
type TicketModelCreate struct {
	config
	mutation *TicketModelMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (tmc *TicketModelCreate) SetUserID(u uuid.UUID) *TicketModelCreate {
	tmc.mutation.SetUserID(u)
	return tmc
}

// SetTenantID sets the "tenant_id" field.
func (tmc *TicketModelCreate) SetTenantID(u uuid.UUID) *TicketModelCreate {
	tmc.mutation.SetTenantID(u)
	return tmc
}

// SetSubject sets the "subject" field.
func (tmc *TicketModelCreate) SetSubject(s string) *TicketModelCreate {
	tmc.mutation.SetSubject(s)
	return tmc
}

// SetBody sets the "body" field.
func (tmc *TicketModelCreate) SetBody(s string) *TicketModelCreate {
	tmc.mutation.SetBody(s)
	return tmc
}

// SetCategory sets the "category" field.
func (tmc *TicketModelCreate) SetCategory(u uuid.UUID) *TicketModelCreate {
	tmc.mutation.SetCategory(u)
	return tmc
}

// SetLabel sets the "label" field.
func (tmc *TicketModelCreate) SetLabel(s string) *TicketModelCreate {
	tmc.mutation.SetLabel(s)
	return tmc
}

// SetAssigneeID sets the "assignee_id" field.
func (tmc *TicketModelCreate) SetAssigneeID(u uuid.UUID) *TicketModelCreate {
	tmc.mutation.SetAssigneeID(u)
	return tmc
}

// SetSeverity sets the "severity" field.
func (tmc *TicketModelCreate) SetSeverity(i int) *TicketModelCreate {
	tmc.mutation.SetSeverity(i)
	return tmc
}

// SetStatus sets the "status" field.
func (tmc *TicketModelCreate) SetStatus(i int) *TicketModelCreate {
	tmc.mutation.SetStatus(i)
	return tmc
}

// SetTime sets the "time" field.
func (tmc *TicketModelCreate) SetTime(t time.Time) *TicketModelCreate {
	tmc.mutation.SetTime(t)
	return tmc
}

// SetID sets the "id" field.
func (tmc *TicketModelCreate) SetID(u uuid.UUID) *TicketModelCreate {
	tmc.mutation.SetID(u)
	return tmc
}

// Mutation returns the TicketModelMutation object of the builder.
func (tmc *TicketModelCreate) Mutation() *TicketModelMutation {
	return tmc.mutation
}

// Save creates the TicketModel in the database.
func (tmc *TicketModelCreate) Save(ctx context.Context) (*TicketModel, error) {
	return withHooks(ctx, tmc.sqlSave, tmc.mutation, tmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tmc *TicketModelCreate) SaveX(ctx context.Context) *TicketModel {
	v, err := tmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmc *TicketModelCreate) Exec(ctx context.Context) error {
	_, err := tmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmc *TicketModelCreate) ExecX(ctx context.Context) {
	if err := tmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tmc *TicketModelCreate) check() error {
	if _, ok := tmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "TicketModel.user_id"`)}
	}
	if _, ok := tmc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "TicketModel.tenant_id"`)}
	}
	if _, ok := tmc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "TicketModel.subject"`)}
	}
	if _, ok := tmc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "TicketModel.body"`)}
	}
	if _, ok := tmc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "TicketModel.category"`)}
	}
	if _, ok := tmc.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "TicketModel.label"`)}
	}
	if _, ok := tmc.mutation.AssigneeID(); !ok {
		return &ValidationError{Name: "assignee_id", err: errors.New(`ent: missing required field "TicketModel.assignee_id"`)}
	}
	if _, ok := tmc.mutation.Severity(); !ok {
		return &ValidationError{Name: "severity", err: errors.New(`ent: missing required field "TicketModel.severity"`)}
	}
	if _, ok := tmc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "TicketModel.status"`)}
	}
	if _, ok := tmc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "TicketModel.time"`)}
	}
	return nil
}

func (tmc *TicketModelCreate) sqlSave(ctx context.Context) (*TicketModel, error) {
	if err := tmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tmc.mutation.id = &_node.ID
	tmc.mutation.done = true
	return _node, nil
}

func (tmc *TicketModelCreate) createSpec() (*TicketModel, *sqlgraph.CreateSpec) {
	var (
		_node = &TicketModel{config: tmc.config}
		_spec = sqlgraph.NewCreateSpec(ticketmodel.Table, sqlgraph.NewFieldSpec(ticketmodel.FieldID, field.TypeUUID))
	)
	if id, ok := tmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tmc.mutation.UserID(); ok {
		_spec.SetField(ticketmodel.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := tmc.mutation.TenantID(); ok {
		_spec.SetField(ticketmodel.FieldTenantID, field.TypeUUID, value)
		_node.TenantID = value
	}
	if value, ok := tmc.mutation.Subject(); ok {
		_spec.SetField(ticketmodel.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := tmc.mutation.Body(); ok {
		_spec.SetField(ticketmodel.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := tmc.mutation.Category(); ok {
		_spec.SetField(ticketmodel.FieldCategory, field.TypeUUID, value)
		_node.Category = value
	}
	if value, ok := tmc.mutation.Label(); ok {
		_spec.SetField(ticketmodel.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := tmc.mutation.AssigneeID(); ok {
		_spec.SetField(ticketmodel.FieldAssigneeID, field.TypeUUID, value)
		_node.AssigneeID = value
	}
	if value, ok := tmc.mutation.Severity(); ok {
		_spec.SetField(ticketmodel.FieldSeverity, field.TypeInt, value)
		_node.Severity = value
	}
	if value, ok := tmc.mutation.Status(); ok {
		_spec.SetField(ticketmodel.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := tmc.mutation.Time(); ok {
		_spec.SetField(ticketmodel.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	return _node, _spec
}

// TicketModelCreateBulk is the builder for creating many TicketModel entities in bulk.
type TicketModelCreateBulk struct {
	config
	builders []*TicketModelCreate
}

// Save creates the TicketModel entities in the database.
func (tmcb *TicketModelCreateBulk) Save(ctx context.Context) ([]*TicketModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmcb.builders))
	nodes := make([]*TicketModel, len(tmcb.builders))
	mutators := make([]Mutator, len(tmcb.builders))
	for i := range tmcb.builders {
		func(i int, root context.Context) {
			builder := tmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketModelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tmcb *TicketModelCreateBulk) SaveX(ctx context.Context) []*TicketModel {
	v, err := tmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmcb *TicketModelCreateBulk) Exec(ctx context.Context) error {
	_, err := tmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcb *TicketModelCreateBulk) ExecX(ctx context.Context) {
	if err := tmcb.Exec(ctx); err != nil {
		panic(err)
	}
}