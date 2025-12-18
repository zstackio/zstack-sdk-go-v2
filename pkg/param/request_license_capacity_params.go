// Copyright (c) ZStack.io, Inc.

package param

// RequestLicenseCapacityDetailParam RequestLicenseCapacity detail param
type RequestLicenseCapacityDetailParam struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	QuotaType string `json:"quotaType" validate:"required"`
	Quota int64 `json:"quota" validate:"required"`
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid" validate:"required"`
	LicenseType string `json:"licenseType" validate:"required"`
}

// RequestLicenseCapacityParam RequestLicenseCapacity request param
type RequestLicenseCapacityParam struct {
	BaseParam
	Params RequestLicenseCapacityDetailParam `json:"params"`
}
