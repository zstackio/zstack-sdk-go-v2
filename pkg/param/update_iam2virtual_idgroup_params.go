// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDGroupDetailParam UpdateIAM2VirtualIDGroup detail param
type UpdateIAM2VirtualIDGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateIAM2VirtualIDGroupParam UpdateIAM2VirtualIDGroup request param
type UpdateIAM2VirtualIDGroupParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupDetailParam `json:"params"`
}
