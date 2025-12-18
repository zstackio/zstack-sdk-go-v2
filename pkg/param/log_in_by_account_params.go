// Copyright (c) ZStack.io, Inc.

package param

// LogInByAccountDetailParam LogInByAccount detail param
type LogInByAccountDetailParam struct {
	AccountName string `json:"accountName" validate:"required"`
	Password string `json:"password" validate:"required"`
	AccountType string `json:"accountType,omitempty"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LogInByAccountParam LogInByAccount request param
type LogInByAccountParam struct {
	BaseParam
	Params LogInByAccountDetailParam `json:"params"`
}
