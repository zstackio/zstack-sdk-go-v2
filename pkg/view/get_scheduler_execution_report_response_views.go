// Copyright (c) ZStack.io, Inc.

package view

// GetSchedulerExecutionReportView GetSchedulerExecutionReport
type GetSchedulerExecutionReportView struct {
	SuccessRecords []int `json:"successRecords,omitempty"`
	FailureRecords []int `json:"failureRecords,omitempty"`
	PartialSuccessRecords []int `json:"partialSuccessRecords,omitempty"`
	WaitingRecords []int `json:"waitingRecords,omitempty"`
	Success bool `json:"success,omitempty"`
}

