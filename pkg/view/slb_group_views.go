// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SlbGroupInventoryView SlbGroup
type SlbGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"backendType,omitempty"`
	rest string `json:"deployType,omitempty"`
	rest string `json:"slbOfferingUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest int64 `json:"configVersion,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []SlbVmInstanceInventoryView `json:"slbVms,omitempty"`
	rest []SlbLoadBalancerInventoryView `json:"lbs,omitempty"`
	rest []SlbGroupL3NetworkRefInventoryView `json:"networks,omitempty"`
	rest []SlbGroupMonitorIpInventoryView `json:"monitorIps,omitempty"`
}

