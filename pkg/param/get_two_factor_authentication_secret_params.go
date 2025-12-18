// Copyright (c) ZStack.io, Inc.

package param

// GetTwoFactorAuthenticationSecretDetailParam GetTwoFactorAuthenticationSecret detail param
type GetTwoFactorAuthenticationSecretDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type string `json:"type" validate:"required"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
}

// GetTwoFactorAuthenticationSecretParam GetTwoFactorAuthenticationSecret request param
type GetTwoFactorAuthenticationSecretParam struct {
	BaseParam
	Params GetTwoFactorAuthenticationSecretDetailParam `json:"params"`
}
