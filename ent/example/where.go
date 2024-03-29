// Code generated by ent, DO NOT EDIT.

package example

import (
	"go-fiber-ent-web-layout/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldName, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldSummary, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldPrice, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Example {
	return predicate.Example(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Example {
	return predicate.Example(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Example {
	return predicate.Example(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Example {
	return predicate.Example(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Example {
	return predicate.Example(sql.FieldContainsFold(FieldName, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Example {
	return predicate.Example(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Example {
	return predicate.Example(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Example {
	return predicate.Example(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Example {
	return predicate.Example(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Example {
	return predicate.Example(sql.FieldContainsFold(FieldSummary, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldPrice, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Example {
	return predicate.Example(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Example {
	return predicate.Example(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Example {
	return predicate.Example(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Example {
	return predicate.Example(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Example {
	return predicate.Example(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Example {
	return predicate.Example(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Example {
	return predicate.Example(sql.FieldNotNull(FieldUpdateTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Example) predicate.Example {
	return predicate.Example(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Example) predicate.Example {
	return predicate.Example(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Example) predicate.Example {
	return predicate.Example(sql.NotPredicates(p))
}
