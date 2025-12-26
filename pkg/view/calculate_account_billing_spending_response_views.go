// Copyright (c) ZStack.io, Inc.

package view

// CalculateAccountBillingSpendingView CalculateAccountBillingSpending
type CalculateAccountBillingSpendingView struct {
	Total float64 `json:"total,omitempty"`
	Spending []SpendingView `json:"spending,omitempty"`
	Success bool `json:"success,omitempty"`
}

