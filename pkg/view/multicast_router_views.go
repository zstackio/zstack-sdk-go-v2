// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MulticastRouterInventoryView MulticastRouter
type MulticastRouterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	RpGroups []MulticastRouterRendezvousPointInventoryView `json:"rpGroups,omitempty"`
	VpcVrs []MulticastRouterVpcVRouterRefInventoryView `json:"vpcVrs,omitempty"`
}

