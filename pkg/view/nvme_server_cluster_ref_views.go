// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NvmeServerClusterRefInventoryView NvmeServerClusterRef
type NvmeServerClusterRefInventoryView struct {
	NvmeServerUuid *string `json:"nvmeServerUuid,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

