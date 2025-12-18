// Copyright (c) ZStack.io, Inc.

package param

// AddUserToGroupDetailParam AddUserToGroup detail param
type AddUserToGroupDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// AddUserToGroupParam AddUserToGroup request param
type AddUserToGroupParam struct {
	BaseParam
	Params AddUserToGroupDetailParam `json:"params"`
}
