// Copyright (c) ZStack.io, Inc.

package view

// CalculateResourceSpendingView CalculateResourceSpending
type CalculateResourceSpendingView struct {
	Spending []ResourceSpendingView `json:"spending,omitempty"`
	Pagination PaginationView `json:"pagination,omitempty"`
	Success bool `json:"success,omitempty"`
}

