// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NvmeTargetInventoryView NvmeTarget
type NvmeTargetInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"nqn,omitempty"`
	rest string `json:"nvmeServerUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest []NvmeLunInventoryView `json:"nvmeLuns,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

