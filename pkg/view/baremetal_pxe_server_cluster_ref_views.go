// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BaremetalPxeServerClusterRefInventoryView BaremetalPxeServerClusterRef
type BaremetalPxeServerClusterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	PxeServerUuid string `json:"pxeServerUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

