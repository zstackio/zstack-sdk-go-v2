// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MulticastRouterInventoryView MulticastRouter
type MulticastRouterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []MulticastRouterRendezvousPointInventoryView `json:"rpGroups,omitempty"`
	rest []MulticastRouterVpcVRouterRefInventoryView `json:"vpcVrs,omitempty"`
}

