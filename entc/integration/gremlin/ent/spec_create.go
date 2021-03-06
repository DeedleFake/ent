// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/spec"
)

// SpecCreate is the builder for creating a Spec entity.
type SpecCreate struct {
	config
	mutation *SpecMutation
	hooks    []Hook
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (sc *SpecCreate) AddCardIDs(ids ...string) *SpecCreate {
	sc.mutation.AddCardIDs(ids...)
	return sc
}

// AddCard adds the "card" edges to the Card entity.
func (sc *SpecCreate) AddCard(c ...*Card) *SpecCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (sc *SpecCreate) Mutation() *SpecMutation {
	return sc.mutation
}

// Save creates the Spec in the database.
func (sc *SpecCreate) Save(ctx context.Context) (*Spec, error) {
	var (
		err  error
		node *Spec
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpecCreate) SaveX(ctx context.Context) *Spec {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sc *SpecCreate) check() error {
	return nil
}

func (sc *SpecCreate) gremlinSave(ctx context.Context) (*Spec, error) {
	res := &gremlin.Response{}
	query, bindings := sc.gremlin().Query()
	if err := sc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	s := &Spec{config: sc.config}
	if err := s.FromResponse(res); err != nil {
		return nil, err
	}
	return s, nil
}

func (sc *SpecCreate) gremlin() *dsl.Traversal {
	v := g.AddV(spec.Label)
	for _, id := range sc.mutation.CardIDs() {
		v.AddE(spec.CardLabel).To(g.V(id)).OutV()
	}
	return v.ValueMap(true)
}

// SpecCreateBulk is the builder for creating many Spec entities in bulk.
type SpecCreateBulk struct {
	config
	builders []*SpecCreate
}
