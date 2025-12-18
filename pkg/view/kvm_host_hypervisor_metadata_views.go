// Copyright (c) ZStack.io, Inc.

package view

import "time"

// KvmHostHypervisorMetadataInventoryView KvmHostHypervisorMetadata
type KvmHostHypervisorMetadataInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"categoryUuid,omitempty"`
	rest string `json:"managementNodeUuid,omitempty"`
	rest string `json:"hypervisor,omitempty"`
	rest string `json:"version,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

