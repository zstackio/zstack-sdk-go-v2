// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2VirtualIDWithLdapDetailParam LoginIAM2VirtualIDWithLdap detail param
type LoginIAM2VirtualIDWithLdapDetailParam struct {
	Uid string `json:"uid" validate:"required"`
	Password string `json:"password" validate:"required"`
	VerifyCode string `json:"verifyCode,omitempty"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2VirtualIDWithLdapParam LoginIAM2VirtualIDWithLdap request param
type LoginIAM2VirtualIDWithLdapParam struct {
	BaseParam
	Params LoginIAM2VirtualIDWithLdapDetailParam `json:"params"`
}
