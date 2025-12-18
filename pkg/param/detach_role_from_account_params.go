// Copyright (c) ZStack.io, Inc.

package param

// DetachRoleFromAccountDetailParam DetachRoleFromAccount detail param
type DetachRoleFromAccountDetailParam struct {
	RoleUuid string `json:"roleUuid" validate:"required"`
	AccountUuid string `json:"accountUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DetachRoleFromAccountParam DetachRoleFromAccount request param
type DetachRoleFromAccountParam struct {
	BaseParam
	Params DetachRoleFromAccountDetailParam `json:"params"`
}
