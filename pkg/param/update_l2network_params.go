// Copyright (c) ZStack.io, Inc.

package param

// UpdateL2NetworkDetailParam UpdateL2Network detail param
type UpdateL2NetworkDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateL2NetworkParam UpdateL2Network request param
type UpdateL2NetworkParam struct {
	BaseParam
	Params UpdateL2NetworkDetailParam `json:"params"`
}
