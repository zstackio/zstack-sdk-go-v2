// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2VirtualIDDetailParam LoginIAM2VirtualID详细参数
type LoginIAM2VirtualIDDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"captchaUuid,omitempty"`
	rest string `json:"verifyCode,omitempty"`
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2VirtualIDParam LoginIAM2VirtualID请求参数
type LoginIAM2VirtualIDParam struct {
	BaseParam
	Params LoginIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

