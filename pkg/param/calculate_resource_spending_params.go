// Copyright (c) ZStack.io, Inc.

package param

// CalculateResourceSpendingDetailParam CalculateResourceSpending详细参数
type CalculateResourceSpendingDetailParam struct {
	rest string `json:"resourceType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"dateStart,omitempty"`
	rest string `json:"dateEnd,omitempty"`
	rest int `json:"start,omitempty"`
	rest int `json:"limit,omitempty"`
}

// CalculateResourceSpendingParam CalculateResourceSpending请求参数
type CalculateResourceSpendingParam struct {
	BaseParam
	Params CalculateResourceSpendingDetailParam `json:"params"` // 详细参数
}

