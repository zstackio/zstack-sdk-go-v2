// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkInterfaceServiceRefInventoryView HostNetworkInterfaceServiceRef
type HostNetworkInterfaceServiceRefInventoryView struct {
	InterfaceUuid string `json:"interfaceUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

