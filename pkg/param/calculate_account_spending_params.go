// Copyright (c) ZStack.io, Inc.

package param

// CalculateAccountSpendingDetailParam CalculateAccountSpending detail param
type CalculateAccountSpendingDetailParam struct {
	AccountUuid string `json:"accountUuid" validate:"required"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	DateStart int64 `json:"dateStart,omitempty"`
	DateEnd int64 `json:"dateEnd,omitempty"`
	Simple bool `json:"simple,omitempty"`
}

// CalculateAccountSpendingParam CalculateAccountSpending request param
type CalculateAccountSpendingParam struct {
	BaseParam
	Params CalculateAccountSpendingDetailParam `json:"params"`
}
