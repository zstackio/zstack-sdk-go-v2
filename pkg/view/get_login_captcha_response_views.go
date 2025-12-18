// Copyright (c) ZStack.io, Inc.

package view

// GetLoginCaptchaView GetLoginCaptcha
type GetLoginCaptchaView struct {
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	Captcha string `json:"captcha,omitempty"`
	Success bool `json:"success,omitempty"`
}

