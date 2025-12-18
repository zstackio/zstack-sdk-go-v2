// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NvmeServerInventoryView NvmeServer
type NvmeServerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"ip,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"transport,omitempty"`
	rest []NvmeTargetInventoryView `json:"nvmeTargets,omitempty"`
	rest []NvmeServerClusterRefInventoryView `json:"nvmeClusterRefs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

