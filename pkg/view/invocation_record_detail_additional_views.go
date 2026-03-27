// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InvocationRecordDetailView InvocationRecordDetail
type InvocationRecordDetailView struct {
	InstanceId string `json:"instanceId,omitempty"`
	InvocationStatus string `json:"invocationStatus,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	FinishTime time.Time `json:"finishTime,omitempty"`
	Output string `json:"output,omitempty"`
}

