// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2VirtualSwitchNetworkInventoryView L2VirtualSwitchNetwork
type L2VirtualSwitchNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	IsDistributed *bool `json:"isDistributed,omitempty"`
	PortGroups []L2PortGroupNetworkInventoryView `json:"portGroups,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	VirtualNetworkId *int `json:"virtualNetworkId,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// QueryL2VirtualSwitchNetworkView QueryL2VirtualSwitchNetwork
type QueryL2VirtualSwitchNetworkView struct {
	Inventories []L2VirtualSwitchNetworkInventoryView `json:"inventories,omitempty"`
}

