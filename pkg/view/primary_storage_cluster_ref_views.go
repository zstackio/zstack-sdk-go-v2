// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PrimaryStorageClusterRefInventoryView PrimaryStorageClusterRef
type PrimaryStorageClusterRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

