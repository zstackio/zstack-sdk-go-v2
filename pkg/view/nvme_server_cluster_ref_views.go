// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NvmeServerClusterRefInventoryView NvmeServerClusterRef
type NvmeServerClusterRefInventoryView struct {
	rest string `json:"nvmeServerUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

