// Copyright (c) ZStack.io, Inc.

package param

// GetOrganizationQuotaUsageDetailParam GetOrganizationQuotaUsage详细参数
type GetOrganizationQuotaUsageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetOrganizationQuotaUsageParam GetOrganizationQuotaUsage请求参数
type GetOrganizationQuotaUsageParam struct {
	BaseParam
	Params GetOrganizationQuotaUsageDetailParam `json:"params"` // 详细参数
}

