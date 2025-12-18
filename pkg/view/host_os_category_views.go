// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostOsCategoryInventoryView HostOsCategory
type HostOsCategoryInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	OsReleaseVersion string `json:"osReleaseVersion,omitempty"`
	MetadataList []KvmHostHypervisorMetadataInventoryView `json:"metadataList,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

