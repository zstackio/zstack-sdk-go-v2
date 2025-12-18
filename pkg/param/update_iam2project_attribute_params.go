// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2ProjectAttributeDetailParam UpdateIAM2ProjectAttribute detail param
type UpdateIAM2ProjectAttributeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// UpdateIAM2ProjectAttributeParam UpdateIAM2ProjectAttribute request param
type UpdateIAM2ProjectAttributeParam struct {
	BaseParam
	Params UpdateIAM2ProjectAttributeDetailParam `json:"params"`
}
