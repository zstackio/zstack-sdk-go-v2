// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HostOsCategoryInventoryView HostOsCategory
type HostOsCategoryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"architecture,omitempty"`
	rest string `json:"osReleaseVersion,omitempty"`
	rest []KvmHostHypervisorMetadataInventoryView `json:"metadataList,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

