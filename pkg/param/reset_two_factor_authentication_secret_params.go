// Copyright (c) ZStack.io, Inc.

package param

// ResetTwoFactorAuthenticationSecretDetailParam ResetTwoFactorAuthenticationSecret详细参数
type ResetTwoFactorAuthenticationSecretDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"captchaUuid,omitempty"`
	rest string `json:"verifyCode,omitempty"`
}

// ResetTwoFactorAuthenticationSecretParam ResetTwoFactorAuthenticationSecret请求参数
type ResetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params ResetTwoFactorAuthenticationSecretDetailParam `json:"params"` // 详细参数
}

