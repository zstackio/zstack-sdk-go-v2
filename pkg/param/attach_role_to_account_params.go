// Copyright (c) ZStack.io, Inc.

package param

// AttachRoleToAccountDetailParam AttachRoleToAccount detail param
type AttachRoleToAccountDetailParam struct {
	RoleUuid string `json:"roleUuid" validate:"required"`
	AccountUuid string `json:"accountUuid" validate:"required"`
}

// AttachRoleToAccountParam AttachRoleToAccount request param
type AttachRoleToAccountParam struct {
	BaseParam
	Params AttachRoleToAccountDetailParam `json:"params"`
}
