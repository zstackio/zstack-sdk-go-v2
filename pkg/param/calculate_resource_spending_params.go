// Copyright (c) ZStack.io, Inc.

package param

// CalculateResourceSpendingDetailParam CalculateResourceSpending detail param
type CalculateResourceSpendingDetailParam struct {
	ResourceType string `json:"resourceType,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	DateEnd string `json:"dateEnd,omitempty"`
	Start int `json:"start,omitempty"`
	Limit int `json:"limit,omitempty"`
}

// CalculateResourceSpendingParam CalculateResourceSpending request param
type CalculateResourceSpendingParam struct {
	BaseParam
	Params CalculateResourceSpendingDetailParam `json:"params"`
}
