// Copyright (c) ZStack.io, Inc.

package param

// RemoveUserFromGroupDetailParam RemoveUserFromGroup详细参数
type RemoveUserFromGroupDetailParam struct {
	rest string `json:"userUuid" validate:"required"` // 必填
	rest string `json:"groupUuid" validate:"required"` // 必填
}

// RemoveUserFromGroupParam RemoveUserFromGroup请求参数
type RemoveUserFromGroupParam struct {
	BaseParam
	Params RemoveUserFromGroupDetailParam `json:"params"` // 详细参数
}

