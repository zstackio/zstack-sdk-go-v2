// Copyright (c) ZStack.io, Inc.

package param

// GetOrganizationQuotaUsageDetailParam GetOrganizationQuotaUsage detail param
type GetOrganizationQuotaUsageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetOrganizationQuotaUsageParam GetOrganizationQuotaUsage request param
type GetOrganizationQuotaUsageParam struct {
	BaseParam
	Params GetOrganizationQuotaUsageDetailParam `json:"params"`
}
