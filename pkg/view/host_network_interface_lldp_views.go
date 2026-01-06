// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkInterfaceLldpInventoryView HostNetworkInterfaceLldp
type HostNetworkInterfaceLldpInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	Mode string `json:"mode,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	NeighborDevice HostNetworkInterfaceLldpRefInventoryView `json:"neighborDevice,omitempty"`
}

// GetHostNetworkInterfaceLldpView GetHostNetworkInterfaceLldp
type GetHostNetworkInterfaceLldpView struct {
	Lldp HostNetworkInterfaceLldpRefInventoryView `json:"lldp,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ChangeHostNetworkInterfaceLldpModeEventView ChangeHostNetworkInterfaceLldpModeEvent
type ChangeHostNetworkInterfaceLldpModeEventView struct {
	Inventories []HostNetworkInterfaceLldpInventoryView `json:"inventories,omitempty"`
}

// QueryHostNetworkInterfaceLldpView QueryHostNetworkInterfaceLldp
type QueryHostNetworkInterfaceLldpView struct {
	Inventories []HostNetworkInterfaceLldpInventoryView `json:"inventories,omitempty"`
}

