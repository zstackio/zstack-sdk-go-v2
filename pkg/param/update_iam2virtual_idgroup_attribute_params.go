// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDGroupAttributeDetailParam UpdateIAM2VirtualIDGroupAttribute detail param
type UpdateIAM2VirtualIDGroupAttributeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2VirtualIDGroupAttributeParam UpdateIAM2VirtualIDGroupAttribute request param
type UpdateIAM2VirtualIDGroupAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDGroupAttributeDetailParam `json:"params"`
}
