// Copyright (c) ZStack.io, Inc.

package param

// GetLoginCaptchaDetailParam GetLoginCaptcha详细参数
type GetLoginCaptchaDetailParam struct {
	rest string `json:"resourceName" validate:"required"` // 必填
	rest string `json:"loginType" validate:"required"` // 必填
	rest string `json:"captchaUuid,omitempty"`
}

// GetLoginCaptchaParam GetLoginCaptcha请求参数
type GetLoginCaptchaParam struct {
	BaseParam
	Params GetLoginCaptchaDetailParam `json:"params"` // 详细参数
}

