// Copyright (c) ZStack.io, Inc.

package view

import "time"

// InstallPathRecycleInventoryView InstallPathRecycle
type InstallPathRecycleInventoryView struct {
	rest int64 `json:"trashId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest string `json:"storageUuid,omitempty"`
	rest string `json:"storageType,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest bool `json:"isFolder,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"trashType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

