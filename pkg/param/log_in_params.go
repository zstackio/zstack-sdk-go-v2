// Copyright (c) ZStack.io, Inc.

package param

// LogInDetailParam LogIn详细参数
type LogInDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
	rest string `json:"captchaUuid,omitempty"`
	rest string `json:"verifyCode,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
	rest map[string]string `json:"properties,omitempty"`
}

// LogInParam LogIn请求参数
type LogInParam struct {
	BaseParam
	Params LogInDetailParam `json:"params"` // 详细参数
}

