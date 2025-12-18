// Copyright (c) ZStack.io, Inc.

package param

// RequestLicenseCapacityDetailParam RequestLicenseCapacity详细参数
type RequestLicenseCapacityDetailParam struct {
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest string `json:"quotaType" validate:"required"` // 必填
	rest int64 `json:"quota" validate:"required"` // 必填
	rest string `json:"clientAuthorizedNodeUuid" validate:"required"` // 必填
	rest string `json:"licenseType" validate:"required"` // 必填
}

// RequestLicenseCapacityParam RequestLicenseCapacity请求参数
type RequestLicenseCapacityParam struct {
	BaseParam
	Params RequestLicenseCapacityDetailParam `json:"params"` // 详细参数
}

