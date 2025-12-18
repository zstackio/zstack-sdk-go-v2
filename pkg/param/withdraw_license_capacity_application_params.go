// Copyright (c) ZStack.io, Inc.

package param

// WithdrawLicenseCapacityApplicationDetailParam WithdrawLicenseCapacityApplication detail param
type WithdrawLicenseCapacityApplicationDetailParam struct {
	ResourceUuidList []string `json:"resourceUuidList" validate:"required"`
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid" validate:"required"`
	LicenseType string `json:"licenseType" validate:"required"`
}

// WithdrawLicenseCapacityApplicationParam WithdrawLicenseCapacityApplication request param
type WithdrawLicenseCapacityApplicationParam struct {
	BaseParam
	Params WithdrawLicenseCapacityApplicationDetailParam `json:"params"`
}
