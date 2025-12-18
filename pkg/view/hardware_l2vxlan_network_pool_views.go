// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HardwareL2VxlanNetworkPoolInventoryView HardwareL2VxlanNetworkPool
type HardwareL2VxlanNetworkPoolInventoryView struct {
	rest string `json:"sdnControllerUuid,omitempty"`
	rest int `json:"startVlan,omitempty"`
	rest int `json:"endVlan,omitempty"`
	rest []VtepInventoryView `json:"attachedVtepRefs,omitempty"`
	rest []RemoteVtepInventoryView `json:"remoteVteps,omitempty"`
	rest []L2VxlanNetworkInventoryView `json:"attachedVxlanNetworkRefs,omitempty"`
	rest []VniRangeInventoryView `json:"attachedVniRanges,omitempty"`
	rest map[string]string `json:"attachedCidrs,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"physicalInterface,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest int `json:"virtualNetworkId,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"attachedClusterUuids,omitempty"`
}

