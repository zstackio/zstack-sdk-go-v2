// Copyright (c) ZStack.io, Inc.

package view

// GetVmInstanceRecoveryPointsView GetVmInstanceRecoveryPoints
type GetVmInstanceRecoveryPointsView struct {
	RecoveryPoints map[string]interface{} `json:"recoveryPoints,omitempty"`
	Success bool `json:"success,omitempty"`
}

