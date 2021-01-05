// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/JustinHaTran/ImageRepo/ent/imagerepo"
	"github.com/JustinHaTran/ImageRepo/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ImageRepoDelete is the builder for deleting a ImageRepo entity.
type ImageRepoDelete struct {
	config
	hooks    []Hook
	mutation *ImageRepoMutation
}

// Where adds a new predicate to the delete builder.
func (ird *ImageRepoDelete) Where(ps ...predicate.ImageRepo) *ImageRepoDelete {
	ird.mutation.predicates = append(ird.mutation.predicates, ps...)
	return ird
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ird *ImageRepoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ird.hooks) == 0 {
		affected, err = ird.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageRepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ird.mutation = mutation
			affected, err = ird.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ird.hooks) - 1; i >= 0; i-- {
			mut = ird.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ird.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ird *ImageRepoDelete) ExecX(ctx context.Context) int {
	n, err := ird.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ird *ImageRepoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: imagerepo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: imagerepo.FieldID,
			},
		},
	}
	if ps := ird.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ird.driver, _spec)
}

// ImageRepoDeleteOne is the builder for deleting a single ImageRepo entity.
type ImageRepoDeleteOne struct {
	ird *ImageRepoDelete
}

// Exec executes the deletion query.
func (irdo *ImageRepoDeleteOne) Exec(ctx context.Context) error {
	n, err := irdo.ird.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{imagerepo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (irdo *ImageRepoDeleteOne) ExecX(ctx context.Context) {
	irdo.ird.ExecX(ctx)
}