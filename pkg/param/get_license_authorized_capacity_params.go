// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseAuthorizedCapacityDetailParam GetLicenseAuthorizedCapacity detail param
type GetLicenseAuthorizedCapacityDetailParam struct {
	ClientAuthorizedNodeUuid string `json:"clientAuthorizedNodeUuid,omitempty"`
	ShowServerCapacity bool `json:"showServerCapacity,omitempty"`
}

// GetLicenseAuthorizedCapacityParam GetLicenseAuthorizedCapacity request param
type GetLicenseAuthorizedCapacityParam struct {
	BaseParam
	Params GetLicenseAuthorizedCapacityDetailParam `json:"params"`
}
