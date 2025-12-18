// Copyright (c) ZStack.io, Inc.

package param

// LogInByLdapDetailParam LogInByLdap详细参数
type LogInByLdapDetailParam struct {
	rest string `json:"uid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"verifyCode,omitempty"`
	rest string `json:"captchaUuid,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LogInByLdapParam LogInByLdap请求参数
type LogInByLdapParam struct {
	BaseParam
	Params LogInByLdapDetailParam `json:"params"` // 详细参数
}

