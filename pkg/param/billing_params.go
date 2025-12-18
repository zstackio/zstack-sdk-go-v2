// Copyright (c) ZStack.io, Inc.

package param

// DeleteBillingDetailParam DeleteBilling详细参数
type DeleteBillingDetailParam struct {
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteBillingParam DeleteBilling请求参数
type DeleteBillingParam struct {
	BaseParam
	Params DeleteBillingDetailParam `json:"params"` // 详细参数
}

