// Copyright (c) ZStack.io, Inc.

package param

// UpgradeToLicenseServerDetailParam UpgradeToLicenseServer详细参数
type UpgradeToLicenseServerDetailParam struct {
}

// UpgradeToLicenseServerParam UpgradeToLicenseServer请求参数
type UpgradeToLicenseServerParam struct {
	BaseParam
	Params UpgradeToLicenseServerDetailParam `json:"params"` // 详细参数
}

