// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PhysicalSwitchPortInventoryView PhysicalSwitchPort
type PhysicalSwitchPortInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"ethTrunkName,omitempty"`
	rest string `json:"portType,omitempty"`
	rest string `json:"peerInterfaceUuid,omitempty"`
	rest string `json:"switchUuid,omitempty"`
	rest string `json:"sdnControllerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

