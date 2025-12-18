// Copyright (c) ZStack.io, Inc.

package view

// MountVmInstanceRecoveryPointEventView MountVmInstanceRecoveryPointEvent
type MountVmInstanceRecoveryPointEventView struct {
	ResourcePath string `json:"resourcePath,omitempty"`
	FailedVolumes map[string]string `json:"failedVolumes,omitempty"`
	Success bool `json:"success,omitempty"`
}

