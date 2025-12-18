// Copyright (c) ZStack.io, Inc.

package param

// UpdateAccessControlListDetailParam UpdateAccessControlList detail param
type UpdateAccessControlListDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAccessControlListParam UpdateAccessControlList request param
type UpdateAccessControlListParam struct {
	BaseParam
	Params UpdateAccessControlListDetailParam `json:"params"`
}
