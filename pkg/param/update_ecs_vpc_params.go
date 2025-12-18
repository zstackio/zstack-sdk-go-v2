// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsVpcDetailParam UpdateEcsVpc detail param
type UpdateEcsVpcDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEcsVpcParam UpdateEcsVpc request param
type UpdateEcsVpcParam struct {
	BaseParam
	Params UpdateEcsVpcDetailParam `json:"params"`
}
