// Code generated by ent, DO NOT EDIT.

package ent

import (
	"uploader/ent/fileentity"
	"uploader/ent/schema"
	"uploader/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileentityFields := schema.FileEntity{}.Fields()
	_ = fileentityFields
	// fileentityDescName is the schema descriptor for name field.
	fileentityDescName := fileentityFields[0].Descriptor()
	// fileentity.NameValidator is a validator for the "name" field. It is called by the builders before save.
	fileentity.NameValidator = fileentityDescName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
