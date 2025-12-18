// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2VirtualIDDetailParam LoginIAM2VirtualID detail param
type LoginIAM2VirtualIDDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	CaptchaUuid string `json:"captchaUuid,omitempty"`
	VerifyCode string `json:"verifyCode,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2VirtualIDParam LoginIAM2VirtualID request param
type LoginIAM2VirtualIDParam struct {
	BaseParam
	Params LoginIAM2VirtualIDDetailParam `json:"params"`
}
