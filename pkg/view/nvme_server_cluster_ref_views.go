// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NvmeServerClusterRefInventoryView NvmeServerClusterRef
type NvmeServerClusterRefInventoryView struct {
	NvmeServerUuid string `json:"nvmeServerUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

