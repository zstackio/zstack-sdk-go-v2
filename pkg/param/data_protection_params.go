// Copyright (c) ZStack.io, Inc.

package param

// StartDataProtectionDetailParam StartDataProtection详细参数
type StartDataProtectionDetailParam struct {
	rest string `json:"encryptType" validate:"required"` // 必填
	rest int `json:"auditsIntegrityDate,omitempty"`
}

// StartDataProtectionParam StartDataProtection请求参数
type StartDataProtectionParam struct {
	BaseParam
	Params StartDataProtectionDetailParam `json:"params"` // 详细参数
}

