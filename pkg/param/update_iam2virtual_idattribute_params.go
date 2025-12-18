// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDAttributeDetailParam UpdateIAM2VirtualIDAttribute detail param
type UpdateIAM2VirtualIDAttributeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2VirtualIDAttributeParam UpdateIAM2VirtualIDAttribute request param
type UpdateIAM2VirtualIDAttributeParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDAttributeDetailParam `json:"params"`
}
