// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kapt/kapt/gql/ent/categorymodel"
	"kapt/kapt/gql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryModelDelete is the builder for deleting a CategoryModel entity.
type CategoryModelDelete struct {
	config
	hooks    []Hook
	mutation *CategoryModelMutation
}

// Where appends a list predicates to the CategoryModelDelete builder.
func (cmd *CategoryModelDelete) Where(ps ...predicate.CategoryModel) *CategoryModelDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *CategoryModelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cmd.sqlExec, cmd.mutation, cmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *CategoryModelDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *CategoryModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(categorymodel.Table, sqlgraph.NewFieldSpec(categorymodel.FieldID, field.TypeUUID))
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cmd.mutation.done = true
	return affected, err
}

// CategoryModelDeleteOne is the builder for deleting a single CategoryModel entity.
type CategoryModelDeleteOne struct {
	cmd *CategoryModelDelete
}

// Where appends a list predicates to the CategoryModelDelete builder.
func (cmdo *CategoryModelDeleteOne) Where(ps ...predicate.CategoryModel) *CategoryModelDeleteOne {
	cmdo.cmd.mutation.Where(ps...)
	return cmdo
}

// Exec executes the deletion query.
func (cmdo *CategoryModelDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{categorymodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *CategoryModelDeleteOne) ExecX(ctx context.Context) {
	if err := cmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
