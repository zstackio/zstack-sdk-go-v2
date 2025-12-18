// Copyright (c) ZStack.io, Inc.

package param

// GetLoginCaptchaDetailParam GetLoginCaptcha detail param
type GetLoginCaptchaDetailParam struct {
	ResourceName string `json:"resourceName" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
}

// GetLoginCaptchaParam GetLoginCaptcha request param
type GetLoginCaptchaParam struct {
	BaseParam
	Params GetLoginCaptchaDetailParam `json:"params"`
}
