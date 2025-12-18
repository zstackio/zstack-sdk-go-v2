// Copyright (c) ZStack.io, Inc.

package param

// RefreshCaptchaDetailParam RefreshCaptcha detail param
type RefreshCaptchaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RefreshCaptchaParam RefreshCaptcha request param
type RefreshCaptchaParam struct {
	BaseParam
	Params RefreshCaptchaDetailParam `json:"params"`
}
