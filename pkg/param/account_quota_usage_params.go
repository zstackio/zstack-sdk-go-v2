// Copyright (c) ZStack.io, Inc.

package param

// GetAccountQuotaUsageDetailParam GetAccountQuotaUsage详细参数
type GetAccountQuotaUsageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetAccountQuotaUsageParam GetAccountQuotaUsage请求参数
type GetAccountQuotaUsageParam struct {
	BaseParam
	Params GetAccountQuotaUsageDetailParam `json:"params"` // 详细参数
}

