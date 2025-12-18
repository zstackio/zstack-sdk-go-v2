// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcHaGroupDetailParam UpdateVpcHaGroup detail param
type UpdateVpcHaGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcHaGroupParam UpdateVpcHaGroup request param
type UpdateVpcHaGroupParam struct {
	BaseParam
	Params UpdateVpcHaGroupDetailParam `json:"params"`
}
