// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2VxlanNetworkPoolInventoryView L2VxlanNetworkPool
type L2VxlanNetworkPoolInventoryView struct {
	AttachedVtepRefs []VtepInventoryView `json:"attachedVtepRefs,omitempty"`
	RemoteVteps []RemoteVtepInventoryView `json:"remoteVteps,omitempty"`
	AttachedVxlanNetworkRefs []L2VxlanNetworkInventoryView `json:"attachedVxlanNetworkRefs,omitempty"`
	AttachedVniRanges []VniRangeInventoryView `json:"attachedVniRanges,omitempty"`
	AttachedCidrs map[string]string `json:"attachedCidrs,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	Isolated bool `json:"isolated,omitempty"`
	Pvlan string `json:"pvlan,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

