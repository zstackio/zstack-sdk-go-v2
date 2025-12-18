// Copyright (c) ZStack.io, Inc.

package param

// ResetTwoFactorAuthenticationSecretDetailParam ResetTwoFactorAuthenticationSecret detail param
type ResetTwoFactorAuthenticationSecretDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type string `json:"type" validate:"required"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
}

// ResetTwoFactorAuthenticationSecretParam ResetTwoFactorAuthenticationSecret request param
type ResetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params ResetTwoFactorAuthenticationSecretDetailParam `json:"params"`
}
