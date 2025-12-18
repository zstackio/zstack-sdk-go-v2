// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VxlanClusterMappingInventoryView VxlanClusterMapping
type VxlanClusterMappingInventoryView struct {
	rest string `json:"vxlanUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest int `json:"vlanId,omitempty"`
	rest string `json:"physicalInterface,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

