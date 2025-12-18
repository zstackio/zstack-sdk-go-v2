// Copyright (c) ZStack.io, Inc.

package param

// UpdateUserDetailParam UpdateUser detail param
type UpdateUserDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	Password string `json:"password,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
}

// UpdateUserParam UpdateUser request param
type UpdateUserParam struct {
	BaseParam
	Params UpdateUserDetailParam `json:"params"`
}
