// Copyright (c) ZStack.io, Inc.

package param

// AttachRoleToAccountDetailParam AttachRoleToAccount详细参数
type AttachRoleToAccountDetailParam struct {
	rest string `json:"roleUuid" validate:"required"` // 必填
	rest string `json:"accountUuid" validate:"required"` // 必填
}

// AttachRoleToAccountParam AttachRoleToAccount请求参数
type AttachRoleToAccountParam struct {
	BaseParam
	Params AttachRoleToAccountDetailParam `json:"params"` // 详细参数
}

