// Copyright (c) ZStack.io, Inc.

package view

// CalculateAccountSpendingView CalculateAccountSpending
type CalculateAccountSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

