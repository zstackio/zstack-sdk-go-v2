// Copyright (c) ZStack.io, Inc.

package param

// LogInDetailParam LogIn detail param
type LogInDetailParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// LogInParam LogIn request param
type LogInParam struct {
	BaseParam
	Params LogInDetailParam `json:"params"`
}
