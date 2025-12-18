// Copyright (c) ZStack.io, Inc.

package param

// CleanupBillingUsageDetailParam CleanupBillingUsage详细参数
type CleanupBillingUsageDetailParam struct {
	rest string `json:"deleteMode,omitempty"`
}

// CleanupBillingUsageParam CleanupBillingUsage请求参数
type CleanupBillingUsageParam struct {
	BaseParam
	Params CleanupBillingUsageDetailParam `json:"params"` // 详细参数
}

