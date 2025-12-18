// Copyright (c) ZStack.io, Inc.

package param

// UpdateEipDetailParam UpdateEip detail param
type UpdateEipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEipParam UpdateEip request param
type UpdateEipParam struct {
	BaseParam
	Params UpdateEipDetailParam `json:"params"`
}
