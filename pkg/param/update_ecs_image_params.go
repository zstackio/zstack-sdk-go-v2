// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsImageDetailParam UpdateEcsImage detail param
type UpdateEcsImageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
}

// UpdateEcsImageParam UpdateEcsImage request param
type UpdateEcsImageParam struct {
	BaseParam
	Params UpdateEcsImageDetailParam `json:"params"`
}
