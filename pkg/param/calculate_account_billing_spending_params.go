// Copyright (c) ZStack.io, Inc.

package param

// CalculateAccountBillingSpendingDetailParam CalculateAccountBillingSpending详细参数
type CalculateAccountBillingSpendingDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest int64 `json:"dateStart,omitempty"`
	rest int64 `json:"dateEnd,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest bool `json:"simple,omitempty"`
}

// CalculateAccountBillingSpendingParam CalculateAccountBillingSpending请求参数
type CalculateAccountBillingSpendingParam struct {
	BaseParam
	Params CalculateAccountBillingSpendingDetailParam `json:"params"` // 详细参数
}

