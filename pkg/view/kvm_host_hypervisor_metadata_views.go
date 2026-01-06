// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// KvmHostHypervisorMetadataInventoryView KvmHostHypervisorMetadata
type KvmHostHypervisorMetadataInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	CategoryUuid string `json:"categoryUuid,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Version string `json:"version,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

