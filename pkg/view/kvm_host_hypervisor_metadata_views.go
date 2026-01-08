// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// KvmHostHypervisorMetadataInventoryView KvmHostHypervisorMetadata
type KvmHostHypervisorMetadataInventoryView struct {
	Uuid               string    `json:"uuid,omitempty"`
	CategoryUuid       string    `json:"categoryUuid,omitempty"`
	ManagementNodeUuid string    `json:"managementNodeUuid,omitempty"`
	Hypervisor         string    `json:"hypervisor,omitempty"`
	Version            string    `json:"version,omitempty"`
	CreateDate         time.Time `json:"createDate,omitempty"`
	LastOpDate         time.Time `json:"lastOpDate,omitempty"`
}
