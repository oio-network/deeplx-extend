// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/oio-network/deeplx-extend/schema/ent/accesslog"
	"github.com/oio-network/deeplx-extend/schema/ent/schema"
	"github.com/oio-network/deeplx-extend/schema/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accesslogMixin := schema.AccessLog{}.Mixin()
	accesslogMixinFields0 := accesslogMixin[0].Fields()
	_ = accesslogMixinFields0
	accesslogFields := schema.AccessLog{}.Fields()
	_ = accesslogFields
	// accesslogDescCreatedAt is the schema descriptor for created_at field.
	accesslogDescCreatedAt := accesslogMixinFields0[0].Descriptor()
	// accesslog.DefaultCreatedAt holds the default value on creation for the created_at field.
	accesslog.DefaultCreatedAt = accesslogDescCreatedAt.Default.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
