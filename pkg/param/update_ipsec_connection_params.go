// Copyright (c) ZStack.io, Inc.

package param

// UpdateIPsecConnectionDetailParam UpdateIPsecConnection detail param
type UpdateIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateIPsecConnectionParam UpdateIPsecConnection request param
type UpdateIPsecConnectionParam struct {
	BaseParam
	Params UpdateIPsecConnectionDetailParam `json:"params"`
}
