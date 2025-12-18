// Copyright (c) ZStack.io, Inc.

package param

// RegisterLicenseServerDetailParam RegisterLicenseServer详细参数
type RegisterLicenseServerDetailParam struct {
	rest string `json:"ip" validate:"required"` // 必填
	rest map[string]interface{} `json:"loginParams" validate:"required"` // 必填
}

// RegisterLicenseServerParam RegisterLicenseServer请求参数
type RegisterLicenseServerParam struct {
	BaseParam
	Params RegisterLicenseServerDetailParam `json:"params"` // 详细参数
}

