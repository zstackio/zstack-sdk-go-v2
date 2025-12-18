// Copyright (c) ZStack.io, Inc.

package param

// PushLicenseAddOnsUsageDetailParam PushLicenseAddOnsUsage详细参数
type PushLicenseAddOnsUsageDetailParam struct {
	rest string `json:"addOnsUsage" validate:"required"` // 必填
}

// PushLicenseAddOnsUsageParam PushLicenseAddOnsUsage请求参数
type PushLicenseAddOnsUsageParam struct {
	BaseParam
	Params PushLicenseAddOnsUsageDetailParam `json:"params"` // 详细参数
}

