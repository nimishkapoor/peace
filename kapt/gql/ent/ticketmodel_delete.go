// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kapt/kapt/gql/ent/predicate"
	"kapt/kapt/gql/ent/ticketmodel"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketModelDelete is the builder for deleting a TicketModel entity.
type TicketModelDelete struct {
	config
	hooks    []Hook
	mutation *TicketModelMutation
}

// Where appends a list predicates to the TicketModelDelete builder.
func (tmd *TicketModelDelete) Where(ps ...predicate.TicketModel) *TicketModelDelete {
	tmd.mutation.Where(ps...)
	return tmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tmd *TicketModelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tmd.sqlExec, tmd.mutation, tmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tmd *TicketModelDelete) ExecX(ctx context.Context) int {
	n, err := tmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tmd *TicketModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ticketmodel.Table, sqlgraph.NewFieldSpec(ticketmodel.FieldID, field.TypeUUID))
	if ps := tmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tmd.mutation.done = true
	return affected, err
}

// TicketModelDeleteOne is the builder for deleting a single TicketModel entity.
type TicketModelDeleteOne struct {
	tmd *TicketModelDelete
}

// Where appends a list predicates to the TicketModelDelete builder.
func (tmdo *TicketModelDeleteOne) Where(ps ...predicate.TicketModel) *TicketModelDeleteOne {
	tmdo.tmd.mutation.Where(ps...)
	return tmdo
}

// Exec executes the deletion query.
func (tmdo *TicketModelDeleteOne) Exec(ctx context.Context) error {
	n, err := tmdo.tmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ticketmodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tmdo *TicketModelDeleteOne) ExecX(ctx context.Context) {
	if err := tmdo.Exec(ctx); err != nil {
		panic(err)
	}
}