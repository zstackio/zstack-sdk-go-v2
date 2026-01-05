// Copyright (c) ZStack.io, Inc.

package param

// GetLicenseNodeUsageDetailsDetailParam GetLicenseNodeUsageDetails detail param
type GetLicenseNodeUsageDetailsDetailParam struct {
	NodeUuid string `json:"nodeUuid,omitempty"`
}

// GetLicenseNodeUsageDetailsParam GetLicenseNodeUsageDetails request param
type GetLicenseNodeUsageDetailsParam struct {
	BaseParam
	Params GetLicenseNodeUsageDetailsDetailParam `json:"params"`
}
