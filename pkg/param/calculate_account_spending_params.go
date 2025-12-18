// Copyright (c) ZStack.io, Inc.

package param

// CalculateAccountSpendingDetailParam CalculateAccountSpending详细参数
type CalculateAccountSpendingDetailParam struct {
	rest string `json:"accountUuid" validate:"required"` // 必填
	rest string `json:"hypervisorType,omitempty"`
	rest int64 `json:"dateStart,omitempty"`
	rest int64 `json:"dateEnd,omitempty"`
	rest bool `json:"simple,omitempty"`
}

// CalculateAccountSpendingParam CalculateAccountSpending请求参数
type CalculateAccountSpendingParam struct {
	BaseParam
	Params CalculateAccountSpendingDetailParam `json:"params"` // 详细参数
}

