// Copyright (c) ZStack.io, Inc.

package param

// DeleteBillingDetailParam DeleteBilling detail param
type DeleteBillingDetailParam struct {
	AccountUuid string `json:"accountUuid,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBillingParam DeleteBilling request param
type DeleteBillingParam struct {
	BaseParam
	Params DeleteBillingDetailParam `json:"params"`
}
