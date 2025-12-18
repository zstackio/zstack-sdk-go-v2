// Copyright (c) ZStack.io, Inc.

package param

// LogInByUserDetailParam LogInByUser详细参数
type LogInByUserDetailParam struct {
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"userName" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LogInByUserParam LogInByUser请求参数
type LogInByUserParam struct {
	BaseParam
	Params LogInByUserDetailParam `json:"params"` // 详细参数
}

