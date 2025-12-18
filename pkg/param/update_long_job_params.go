// Copyright (c) ZStack.io, Inc.

package param

// UpdateLongJobDetailParam UpdateLongJob detail param
type UpdateLongJobDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateLongJobParam UpdateLongJob request param
type UpdateLongJobParam struct {
	BaseParam
	Params UpdateLongJobDetailParam `json:"params"`
}
