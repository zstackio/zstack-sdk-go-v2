// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ClusterInventoryView Cluster
type ClusterInventoryView struct {
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"hypervisorType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"architecture,omitempty"`
}

