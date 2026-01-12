// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ProvisionNetworkInventoryView BareMetal2ProvisionNetwork
type BareMetal2ProvisionNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	Description *string `json:"description,omitempty"`
	DhcpInterface *string `json:"dhcpInterface,omitempty"`
	DhcpRangeStartIp *string `json:"dhcpRangeStartIp,omitempty"`
	DhcpRangeEndIp *string `json:"dhcpRangeEndIp,omitempty"`
	DhcpRangeNetmask *string `json:"dhcpRangeNetmask,omitempty"`
	DhcpRangeGateway *string `json:"dhcpRangeGateway,omitempty"`
	DhcpRangeNetworkCidr *string `json:"dhcpRangeNetworkCidr,omitempty"`
	State *string `json:"state,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// UpdateBareMetal2ProvisionNetworkEventView UpdateBareMetal2ProvisionNetworkEvent
type UpdateBareMetal2ProvisionNetworkEventView struct {
	Inventory BareMetal2ProvisionNetworkInventoryView `json:"inventory,omitempty"`
}

// QueryBareMetal2ProvisionNetworkView QueryBareMetal2ProvisionNetwork
type QueryBareMetal2ProvisionNetworkView struct {
	Inventories []BareMetal2ProvisionNetworkInventoryView `json:"inventories,omitempty"`
}

// AttachBareMetal2ProvisionNetworkToClusterEventView AttachBareMetal2ProvisionNetworkToClusterEvent
type AttachBareMetal2ProvisionNetworkToClusterEventView struct {
	Inventory BareMetal2ProvisionNetworkInventoryView `json:"inventory,omitempty"`
}

// CreateBareMetal2ProvisionNetworkEventView CreateBareMetal2ProvisionNetworkEvent
type CreateBareMetal2ProvisionNetworkEventView struct {
	Inventory BareMetal2ProvisionNetworkInventoryView `json:"inventory,omitempty"`
}

// DeleteBareMetal2ProvisionNetworkEventView DeleteBareMetal2ProvisionNetworkEvent
type DeleteBareMetal2ProvisionNetworkEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachBareMetal2ProvisionNetworkFromClusterEventView DetachBareMetal2ProvisionNetworkFromClusterEvent
type DetachBareMetal2ProvisionNetworkFromClusterEventView struct {
	Inventory BareMetal2ProvisionNetworkInventoryView `json:"inventory,omitempty"`
}

// ChangeBareMetal2ProvisionNetworkStateEventView ChangeBareMetal2ProvisionNetworkStateEvent
type ChangeBareMetal2ProvisionNetworkStateEventView struct {
	Inventory BareMetal2ProvisionNetworkInventoryView `json:"inventory,omitempty"`
}

