// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/emstockbasicinfo"
	"github.com/suyuan32/simple-admin-efinance-rpc/ent/predicate"
)

// EmStockBasicInfoDelete is the builder for deleting a EmStockBasicInfo entity.
type EmStockBasicInfoDelete struct {
	config
	hooks    []Hook
	mutation *EmStockBasicInfoMutation
}

// Where appends a list predicates to the EmStockBasicInfoDelete builder.
func (esbid *EmStockBasicInfoDelete) Where(ps ...predicate.EmStockBasicInfo) *EmStockBasicInfoDelete {
	esbid.mutation.Where(ps...)
	return esbid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (esbid *EmStockBasicInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, esbid.sqlExec, esbid.mutation, esbid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (esbid *EmStockBasicInfoDelete) ExecX(ctx context.Context) int {
	n, err := esbid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (esbid *EmStockBasicInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emstockbasicinfo.Table, sqlgraph.NewFieldSpec(emstockbasicinfo.FieldID, field.TypeInt32))
	if ps := esbid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, esbid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	esbid.mutation.done = true
	return affected, err
}

// EmStockBasicInfoDeleteOne is the builder for deleting a single EmStockBasicInfo entity.
type EmStockBasicInfoDeleteOne struct {
	esbid *EmStockBasicInfoDelete
}

// Where appends a list predicates to the EmStockBasicInfoDelete builder.
func (esbido *EmStockBasicInfoDeleteOne) Where(ps ...predicate.EmStockBasicInfo) *EmStockBasicInfoDeleteOne {
	esbido.esbid.mutation.Where(ps...)
	return esbido
}

// Exec executes the deletion query.
func (esbido *EmStockBasicInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := esbido.esbid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emstockbasicinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (esbido *EmStockBasicInfoDeleteOne) ExecX(ctx context.Context) {
	if err := esbido.Exec(ctx); err != nil {
		panic(err)
	}
}
