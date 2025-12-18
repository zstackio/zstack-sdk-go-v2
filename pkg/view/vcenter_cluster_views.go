// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VCenterClusterInventoryView VCenterCluster
type VCenterClusterInventoryView struct {
	rest string `json:"vCenterUuid,omitempty"`
	rest string `json:"morval,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
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

