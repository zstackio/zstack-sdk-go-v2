// Copyright (c) ZStack.io, Inc.

package param

// RefreshCaptchaDetailParam RefreshCaptcha详细参数
type RefreshCaptchaDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// RefreshCaptchaParam RefreshCaptcha请求参数
type RefreshCaptchaParam struct {
	BaseParam
	Params RefreshCaptchaDetailParam `json:"params"` // 详细参数
}

