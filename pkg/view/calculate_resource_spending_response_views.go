// Copyright (c) ZStack.io, Inc.

package view

// CalculateResourceSpendingView CalculateResourceSpending
type CalculateResourceSpendingView struct {
	Spending []interface{} `json:"spending,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Success bool `json:"success,omitempty"`
}

