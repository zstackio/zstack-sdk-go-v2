// Copyright (c) ZStack.io, Inc.

package view

// DescribeVmInstanceRecoveryPointView DescribeVmInstanceRecoveryPoint
type DescribeVmInstanceRecoveryPointView struct {
	RealSizes map[string]int64 `json:"realSizes,omitempty"`
	VirtualSizes map[string]int64 `json:"virtualSizes,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Success bool `json:"success,omitempty"`
}

