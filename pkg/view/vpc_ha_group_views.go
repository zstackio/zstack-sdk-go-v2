// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcHaGroupInventoryView VpcHaGroup
type VpcHaGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest []VpcHaGroupMonitorIpInventoryView `json:"monitors,omitempty"`
	rest []VpcHaGroupApplianceVmRefInventoryView `json:"vrRefs,omitempty"`
	rest []VpcHaGroupNetworkServiceRefInventoryView `json:"services,omitempty"`
	rest []VpcHaGroupVipRefInventoryView `json:"usedIps,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

