// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package car

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/customid/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BeforeID applies equality check predicate on the "before_id" field. It's identical to BeforeIDEQ.
func BeforeID(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBeforeID), v))
	})
}

// AfterID applies equality check predicate on the "after_id" field. It's identical to AfterIDEQ.
func AfterID(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAfterID), v))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// BeforeIDEQ applies the EQ predicate on the "before_id" field.
func BeforeIDEQ(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBeforeID), v))
	})
}

// BeforeIDNEQ applies the NEQ predicate on the "before_id" field.
func BeforeIDNEQ(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBeforeID), v))
	})
}

// BeforeIDIn applies the In predicate on the "before_id" field.
func BeforeIDIn(vs ...float64) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBeforeID), v...))
	})
}

// BeforeIDNotIn applies the NotIn predicate on the "before_id" field.
func BeforeIDNotIn(vs ...float64) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBeforeID), v...))
	})
}

// BeforeIDGT applies the GT predicate on the "before_id" field.
func BeforeIDGT(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBeforeID), v))
	})
}

// BeforeIDGTE applies the GTE predicate on the "before_id" field.
func BeforeIDGTE(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBeforeID), v))
	})
}

// BeforeIDLT applies the LT predicate on the "before_id" field.
func BeforeIDLT(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBeforeID), v))
	})
}

// BeforeIDLTE applies the LTE predicate on the "before_id" field.
func BeforeIDLTE(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBeforeID), v))
	})
}

// BeforeIDIsNil applies the IsNil predicate on the "before_id" field.
func BeforeIDIsNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBeforeID)))
	})
}

// BeforeIDNotNil applies the NotNil predicate on the "before_id" field.
func BeforeIDNotNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBeforeID)))
	})
}

// AfterIDEQ applies the EQ predicate on the "after_id" field.
func AfterIDEQ(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAfterID), v))
	})
}

// AfterIDNEQ applies the NEQ predicate on the "after_id" field.
func AfterIDNEQ(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAfterID), v))
	})
}

// AfterIDIn applies the In predicate on the "after_id" field.
func AfterIDIn(vs ...float64) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAfterID), v...))
	})
}

// AfterIDNotIn applies the NotIn predicate on the "after_id" field.
func AfterIDNotIn(vs ...float64) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAfterID), v...))
	})
}

// AfterIDGT applies the GT predicate on the "after_id" field.
func AfterIDGT(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAfterID), v))
	})
}

// AfterIDGTE applies the GTE predicate on the "after_id" field.
func AfterIDGTE(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAfterID), v))
	})
}

// AfterIDLT applies the LT predicate on the "after_id" field.
func AfterIDLT(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAfterID), v))
	})
}

// AfterIDLTE applies the LTE predicate on the "after_id" field.
func AfterIDLTE(v float64) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAfterID), v))
	})
}

// AfterIDIsNil applies the IsNil predicate on the "after_id" field.
func AfterIDIsNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAfterID)))
	})
}

// AfterIDNotNil applies the NotNil predicate on the "after_id" field.
func AfterIDNotNil() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAfterID)))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModel), v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModel), v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModel), v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModel), v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModel), v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModel), v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModel), v))
	})
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModel), v))
	})
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModel), v))
	})
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModel), v))
	})
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModel), v))
	})
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModel), v))
	})
}

// OwnerID applies the EQ predicate on the OwnerColumn field.
func OwnerID(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(OwnerColumn), id))
	})
}

// OwnerIDEQ applies the EQ predicate on the pet_cars field.
func OwnerIDEQ(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(OwnerColumn), id))
	})
}

// OwnerIDNEQ applies the NEQ predicate on the pet_cars field.
func OwnerIDNEQ(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(OwnerColumn), id))
	})
}

// OwnerIDIn applies the In predicate on the pet_cars field.
func OwnerIDIn(ids ...string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(OwnerColumn), v...))
	})
}

// OwnerIDNotIn applies the NotIn predicate on the pet_cars field.
func OwnerIDNotIn(ids ...string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(OwnerColumn), v...))
	})
}

// OwnerIDGT applies the GT predicate on the pet_cars field.
func OwnerIDGT(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(OwnerColumn), id))
	})
}

// OwnerIDGTE applies the GTE predicate on the pet_cars field.
func OwnerIDGTE(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(OwnerColumn), id))
	})
}

// OwnerIDLT applies the LT predicate on the pet_cars field.
func OwnerIDLT(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(OwnerColumn), id))
	})
}

// OwnerIDLTE applies the LTE predicate on the pet_cars field.
func OwnerIDLTE(id string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(OwnerColumn), id))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Pet) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
