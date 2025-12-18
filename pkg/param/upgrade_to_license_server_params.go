// Copyright (c) ZStack.io, Inc.

package param

// UpgradeToLicenseServerDetailParam UpgradeToLicenseServer detail param
type UpgradeToLicenseServerDetailParam struct {
}

// UpgradeToLicenseServerParam UpgradeToLicenseServer request param
type UpgradeToLicenseServerParam struct {
	BaseParam
	Params UpgradeToLicenseServerDetailParam `json:"params"`
}
