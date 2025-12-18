// Copyright (c) ZStack.io, Inc.

package param

// LogInByAccountDetailParam LogInByAccount详细参数
type LogInByAccountDetailParam struct {
	rest string `json:"accountName" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"accountType,omitempty"`
	rest string `json:"captchaUuid,omitempty"`
	rest string `json:"verifyCode,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LogInByAccountParam LogInByAccount请求参数
type LogInByAccountParam struct {
	BaseParam
	Params LogInByAccountDetailParam `json:"params"` // 详细参数
}

