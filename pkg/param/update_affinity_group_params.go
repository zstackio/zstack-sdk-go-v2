// Copyright (c) ZStack.io, Inc.

package param

// UpdateAffinityGroupDetailParam UpdateAffinityGroup detail param
type UpdateAffinityGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAffinityGroupParam UpdateAffinityGroup request param
type UpdateAffinityGroupParam struct {
	BaseParam
	Params UpdateAffinityGroupDetailParam `json:"params"`
}
