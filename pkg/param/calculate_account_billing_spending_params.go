// Copyright (c) ZStack.io, Inc.

package param

// CalculateAccountBillingSpendingDetailParam CalculateAccountBillingSpending detail param
type CalculateAccountBillingSpendingDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	DateStart int64 `json:"dateStart,omitempty"`
	DateEnd int64 `json:"dateEnd,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Simple bool `json:"simple,omitempty"`
}

// CalculateAccountBillingSpendingParam CalculateAccountBillingSpending request param
type CalculateAccountBillingSpendingParam struct {
	BaseParam
	Params CalculateAccountBillingSpendingDetailParam `json:"params"`
}
