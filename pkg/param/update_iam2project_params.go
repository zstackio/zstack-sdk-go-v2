// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2ProjectDetailParam UpdateIAM2Project detail param
type UpdateIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateIAM2ProjectParam UpdateIAM2Project request param
type UpdateIAM2ProjectParam struct {
	BaseParam
	Params UpdateIAM2ProjectDetailParam `json:"params"`
}
