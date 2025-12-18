// Copyright (c) ZStack.io, Inc.

package param

// UpdateHybridEipDetailParam UpdateHybridEip detail param
type UpdateHybridEipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
}

// UpdateHybridEipParam UpdateHybridEip request param
type UpdateHybridEipParam struct {
	BaseParam
	Params UpdateHybridEipDetailParam `json:"params"`
}
