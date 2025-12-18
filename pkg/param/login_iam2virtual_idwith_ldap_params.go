// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2VirtualIDWithLdapDetailParam LoginIAM2VirtualIDWithLdap详细参数
type LoginIAM2VirtualIDWithLdapDetailParam struct {
	rest string `json:"uid" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"verifyCode,omitempty"`
	rest string `json:"captchaUuid,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2VirtualIDWithLdapParam LoginIAM2VirtualIDWithLdap请求参数
type LoginIAM2VirtualIDWithLdapParam struct {
	BaseParam
	Params LoginIAM2VirtualIDWithLdapDetailParam `json:"params"` // 详细参数
}

