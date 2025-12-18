// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IscsiServerClusterRefInventoryView IscsiServerClusterRef
type IscsiServerClusterRefInventoryView struct {
	rest string `json:"iscsiServerUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

