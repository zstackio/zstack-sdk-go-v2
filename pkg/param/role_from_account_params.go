// Copyright (c) ZStack.io, Inc.

package param

// DetachRoleFromAccountDetailParam DetachRoleFromAccount详细参数
type DetachRoleFromAccountDetailParam struct {
	rest string `json:"roleUuid" validate:"required"` // 必填
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DetachRoleFromAccountParam DetachRoleFromAccount请求参数
type DetachRoleFromAccountParam struct {
	BaseParam
	Params DetachRoleFromAccountDetailParam `json:"params"` // 详细参数
}

