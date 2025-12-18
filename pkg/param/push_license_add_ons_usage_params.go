// Copyright (c) ZStack.io, Inc.

package param

// PushLicenseAddOnsUsageDetailParam PushLicenseAddOnsUsage detail param
type PushLicenseAddOnsUsageDetailParam struct {
	AddOnsUsage string `json:"addOnsUsage" validate:"required"`
}

// PushLicenseAddOnsUsageParam PushLicenseAddOnsUsage request param
type PushLicenseAddOnsUsageParam struct {
	BaseParam
	Params PushLicenseAddOnsUsageDetailParam `json:"params"`
}
