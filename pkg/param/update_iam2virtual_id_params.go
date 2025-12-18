// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2VirtualIDDetailParam UpdateIAM2VirtualID detail param
type UpdateIAM2VirtualIDDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Password string `json:"password,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
}

// UpdateIAM2VirtualIDParam UpdateIAM2VirtualID request param
type UpdateIAM2VirtualIDParam struct {
	BaseParam
	Params UpdateIAM2VirtualIDDetailParam `json:"params"`
}
