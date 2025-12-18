// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IscsiServerClusterRefInventoryView IscsiServerClusterRef
type IscsiServerClusterRefInventoryView struct {
	IscsiServerUuid string `json:"iscsiServerUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

