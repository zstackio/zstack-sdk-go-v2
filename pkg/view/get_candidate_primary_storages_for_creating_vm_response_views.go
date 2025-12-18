// Copyright (c) ZStack.io, Inc.

package view

// GetCandidatePrimaryStoragesForCreatingVmView GetCandidatePrimaryStoragesForCreatingVm
type GetCandidatePrimaryStoragesForCreatingVmView struct {
	RootVolumePrimaryStorages []PrimaryStorageInventoryView `json:"rootVolumePrimaryStorages,omitempty"`
	DataVolumePrimaryStorages map[string]interface{} `json:"dataVolumePrimaryStorages,omitempty"`
	Success bool `json:"success,omitempty"`
}

