// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-fiber-ent-web-layout/ent/example"
	"go-fiber-ent-web-layout/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	exampleFields := schema.Example{}.Fields()
	_ = exampleFields
	// exampleDescName is the schema descriptor for name field.
	exampleDescName := exampleFields[1].Descriptor()
	// example.NameValidator is a validator for the "name" field. It is called by the builders before save.
	example.NameValidator = exampleDescName.Validators[0].(func(string) error)
	// exampleDescPrice is the schema descriptor for price field.
	exampleDescPrice := exampleFields[3].Descriptor()
	// example.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	example.PriceValidator = exampleDescPrice.Validators[0].(func(float64) error)
	// exampleDescID is the schema descriptor for id field.
	exampleDescID := exampleFields[0].Descriptor()
	// example.IDValidator is a validator for the "id" field. It is called by the builders before save.
	example.IDValidator = exampleDescID.Validators[0].(func(int) error)
}
