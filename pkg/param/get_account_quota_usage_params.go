// Copyright (c) ZStack.io, Inc.

package param

// GetAccountQuotaUsageDetailParam GetAccountQuotaUsage detail param
type GetAccountQuotaUsageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetAccountQuotaUsageParam GetAccountQuotaUsage request param
type GetAccountQuotaUsageParam struct {
	BaseParam
	Params GetAccountQuotaUsageDetailParam `json:"params"`
}
