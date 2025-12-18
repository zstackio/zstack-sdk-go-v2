// Copyright (c) ZStack.io, Inc.

package param

// UpdateTagDetailParam UpdateTag detail param
type UpdateTagDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Color string `json:"color,omitempty"`
}

// UpdateTagParam UpdateTag request param
type UpdateTagParam struct {
	BaseParam
	Params UpdateTagDetailParam `json:"params"`
}
