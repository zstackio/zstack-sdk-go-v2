// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcHaGroupInventoryView VpcHaGroup
type VpcHaGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Monitors []VpcHaGroupMonitorIpInventoryView `json:"monitors,omitempty"`
	VrRefs []VpcHaGroupApplianceVmRefInventoryView `json:"vrRefs,omitempty"`
	Services []VpcHaGroupNetworkServiceRefInventoryView `json:"services,omitempty"`
	UsedIps []VpcHaGroupVipRefInventoryView `json:"usedIps,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

