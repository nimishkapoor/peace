// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kapt/kapt/gql/ent/threadmodel"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ThreadModelCreate is the builder for creating a ThreadModel entity.
type ThreadModelCreate struct {
	config
	mutation *ThreadModelMutation
	hooks    []Hook
}

// SetBody sets the "body" field.
func (tmc *ThreadModelCreate) SetBody(s string) *ThreadModelCreate {
	tmc.mutation.SetBody(s)
	return tmc
}

// SetStatus sets the "status" field.
func (tmc *ThreadModelCreate) SetStatus(i int) *ThreadModelCreate {
	tmc.mutation.SetStatus(i)
	return tmc
}

// SetID sets the "id" field.
func (tmc *ThreadModelCreate) SetID(u uuid.UUID) *ThreadModelCreate {
	tmc.mutation.SetID(u)
	return tmc
}

// Mutation returns the ThreadModelMutation object of the builder.
func (tmc *ThreadModelCreate) Mutation() *ThreadModelMutation {
	return tmc.mutation
}

// Save creates the ThreadModel in the database.
func (tmc *ThreadModelCreate) Save(ctx context.Context) (*ThreadModel, error) {
	return withHooks(ctx, tmc.sqlSave, tmc.mutation, tmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tmc *ThreadModelCreate) SaveX(ctx context.Context) *ThreadModel {
	v, err := tmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmc *ThreadModelCreate) Exec(ctx context.Context) error {
	_, err := tmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmc *ThreadModelCreate) ExecX(ctx context.Context) {
	if err := tmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tmc *ThreadModelCreate) check() error {
	if _, ok := tmc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "ThreadModel.body"`)}
	}
	if _, ok := tmc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ThreadModel.status"`)}
	}
	return nil
}

func (tmc *ThreadModelCreate) sqlSave(ctx context.Context) (*ThreadModel, error) {
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

func (tmc *ThreadModelCreate) createSpec() (*ThreadModel, *sqlgraph.CreateSpec) {
	var (
		_node = &ThreadModel{config: tmc.config}
		_spec = sqlgraph.NewCreateSpec(threadmodel.Table, sqlgraph.NewFieldSpec(threadmodel.FieldID, field.TypeUUID))
	)
	if id, ok := tmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tmc.mutation.Body(); ok {
		_spec.SetField(threadmodel.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := tmc.mutation.Status(); ok {
		_spec.SetField(threadmodel.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// ThreadModelCreateBulk is the builder for creating many ThreadModel entities in bulk.
type ThreadModelCreateBulk struct {
	config
	builders []*ThreadModelCreate
}

// Save creates the ThreadModel entities in the database.
func (tmcb *ThreadModelCreateBulk) Save(ctx context.Context) ([]*ThreadModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmcb.builders))
	nodes := make([]*ThreadModel, len(tmcb.builders))
	mutators := make([]Mutator, len(tmcb.builders))
	for i := range tmcb.builders {
		func(i int, root context.Context) {
			builder := tmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadModelMutation)
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
func (tmcb *ThreadModelCreateBulk) SaveX(ctx context.Context) []*ThreadModel {
	v, err := tmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmcb *ThreadModelCreateBulk) Exec(ctx context.Context) error {
	_, err := tmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcb *ThreadModelCreateBulk) ExecX(ctx context.Context) {
	if err := tmcb.Exec(ctx); err != nil {
		panic(err)
	}
}