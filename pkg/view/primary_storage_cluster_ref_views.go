// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PrimaryStorageClusterRefInventoryView PrimaryStorageClusterRef
type PrimaryStorageClusterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

