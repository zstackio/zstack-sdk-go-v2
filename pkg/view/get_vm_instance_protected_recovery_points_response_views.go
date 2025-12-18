// Copyright (c) ZStack.io, Inc.

package view

// GetVmInstanceProtectedRecoveryPointsView GetVmInstanceProtectedRecoveryPoints
type GetVmInstanceProtectedRecoveryPointsView struct {
	RecoveryPoints map[string]interface{} `json:"recoveryPoints,omitempty"`
	Success bool `json:"success,omitempty"`
}

