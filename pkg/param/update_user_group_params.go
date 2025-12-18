// Copyright (c) ZStack.io, Inc.

package param

// UpdateUserGroupDetailParam UpdateUserGroup detail param
type UpdateUserGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateUserGroupParam UpdateUserGroup request param
type UpdateUserGroupParam struct {
	BaseParam
	Params UpdateUserGroupDetailParam `json:"params"`
}
