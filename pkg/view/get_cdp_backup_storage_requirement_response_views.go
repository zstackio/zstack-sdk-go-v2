// Copyright (c) ZStack.io, Inc.

package view

// GetCdpBackupStorageRequirementView GetCdpBackupStorageRequirement
type GetCdpBackupStorageRequirementView struct {
	NextStep string `json:"nextStep,omitempty"`
	Required map[string]string `json:"required,omitempty"`
	Current map[string]string `json:"current,omitempty"`
	Success bool `json:"success,omitempty"`
}

