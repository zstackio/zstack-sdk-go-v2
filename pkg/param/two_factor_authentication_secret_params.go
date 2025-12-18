// Copyright (c) ZStack.io, Inc.

package param

// GetTwoFactorAuthenticationSecretDetailParam GetTwoFactorAuthenticationSecret详细参数
type GetTwoFactorAuthenticationSecretDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"captchaUuid,omitempty"`
	rest string `json:"verifyCode,omitempty"`
}

// GetTwoFactorAuthenticationSecretParam GetTwoFactorAuthenticationSecret请求参数
type GetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationSecretDetailParam `json:"params"` // 详细参数
}

