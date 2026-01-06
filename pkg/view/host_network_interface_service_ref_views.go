// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkInterfaceServiceRefInventoryView HostNetworkInterfaceServiceRef
type HostNetworkInterfaceServiceRefInventoryView struct {
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// SetServiceTypeOnHostNetworkInterfaceEventView SetServiceTypeOnHostNetworkInterfaceEvent
type SetServiceTypeOnHostNetworkInterfaceEventView struct {
	Inventory []HostNetworkInterfaceServiceRefInventoryView `json:"inventory,omitempty"`
}

