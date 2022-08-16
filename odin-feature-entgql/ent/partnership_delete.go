// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"odin/ent/partnership"
	"odin/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PartnershipDelete is the builder for deleting a Partnership entity.
type PartnershipDelete struct {
	config
	hooks    []Hook
	mutation *PartnershipMutation
}

// Where appends a list predicates to the PartnershipDelete builder.
func (pd *PartnershipDelete) Where(ps ...predicate.Partnership) *PartnershipDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PartnershipDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartnershipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PartnershipDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PartnershipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: partnership.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partnership.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PartnershipDeleteOne is the builder for deleting a single Partnership entity.
type PartnershipDeleteOne struct {
	pd *PartnershipDelete
}

// Exec executes the deletion query.
func (pdo *PartnershipDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{partnership.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PartnershipDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}