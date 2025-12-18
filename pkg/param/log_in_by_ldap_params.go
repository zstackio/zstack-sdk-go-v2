// Copyright (c) ZStack.io, Inc.

package param

// LogInByLdapDetailParam LogInByLdap detail param
type LogInByLdapDetailParam struct {
	Uid string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
	VerifyCode string `json:"verifyCode,omitempty"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LogInByLdapParam LogInByLdap request param
type LogInByLdapParam struct {
	BaseParam
	Params LogInByLdapDetailParam `json:"params"`
}
