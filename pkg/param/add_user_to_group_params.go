// Copyright (c) ZStack.io, Inc.

package param

// AddUserToGroupDetailParam AddUserToGroup详细参数
type AddUserToGroupDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// AddUserToGroupParam AddUserToGroup请求参数
type AddUserToGroupParam struct {
	BaseParam
	Params AddUserToGroupDetailParam `json:"params"` // 详细参数
}

