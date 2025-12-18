// Copyright (c) ZStack.io, Inc.

package param

// UpdateAccountDetailParam UpdateAccount detail param
type UpdateAccountDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Password string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
}

// UpdateAccountParam UpdateAccount request param
type UpdateAccountParam struct {
	BaseParam
	Params UpdateAccountDetailParam `json:"params"`
}
