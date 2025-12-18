// Copyright (c) ZStack.io, Inc.

package param

// WithdrawLicenseCapacityApplicationDetailParam WithdrawLicenseCapacityApplication详细参数
type WithdrawLicenseCapacityApplicationDetailParam struct {
	rest []string `json:"resourceUuidList" validate:"required"` // 必填
	rest string `json:"clientAuthorizedNodeUuid" validate:"required"` // 必填
	rest string `json:"licenseType" validate:"required"` // 必填
}

// WithdrawLicenseCapacityApplicationParam WithdrawLicenseCapacityApplication请求参数
type WithdrawLicenseCapacityApplicationParam struct {
	BaseParam
	Params WithdrawLicenseCapacityApplicationDetailParam `json:"params"` // 详细参数
}

