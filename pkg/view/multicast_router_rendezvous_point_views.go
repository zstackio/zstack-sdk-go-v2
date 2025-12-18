// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MulticastRouterRendezvousPointInventoryView MulticastRouterRendezvousPoint
type MulticastRouterRendezvousPointInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"multicastRouterUuid,omitempty"`
	rest string `json:"rpAddress,omitempty"`
	rest string `json:"groupAddress,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

