// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PhysicalSwitchPortInventoryView PhysicalSwitchPort
type PhysicalSwitchPortInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	EthTrunkName string `json:"ethTrunkName,omitempty"`
	PortType string `json:"portType,omitempty"`
	PeerInterfaceUuid string `json:"peerInterfaceUuid,omitempty"`
	SwitchUuid string `json:"switchUuid,omitempty"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

