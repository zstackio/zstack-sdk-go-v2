// Copyright (c) ZStack.io, Inc.

package param

// CleanupBillingUsageDetailParam CleanupBillingUsage detail param
type CleanupBillingUsageDetailParam struct {
	DeleteMode string `json:"deleteMode,omitempty"`
}

// CleanupBillingUsageParam CleanupBillingUsage request param
type CleanupBillingUsageParam struct {
	BaseParam
	Params CleanupBillingUsageDetailParam `json:"params"`
}
