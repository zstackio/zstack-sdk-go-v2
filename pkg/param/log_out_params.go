// Copyright (c) ZStack.io, Inc.

package param

// LogOutDetailParam LogOut详细参数
type LogOutDetailParam struct {
	rest string `json:"sessionUuid,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LogOutParam LogOut请求参数
type LogOutParam struct {
	BaseParam
	Params LogOutDetailParam `json:"params"` // 详细参数
}

