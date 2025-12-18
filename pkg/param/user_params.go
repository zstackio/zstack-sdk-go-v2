// Copyright (c) ZStack.io, Inc.

package param

// DeleteUserDetailParam DeleteUser详细参数
type DeleteUserDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteUserParam DeleteUser请求参数
type DeleteUserParam struct {
	BaseParam
	Params DeleteUserDetailParam `json:"params"` // 详细参数
}

