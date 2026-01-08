// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MulticastRouterRendezvousPointInventoryView MulticastRouterRendezvousPoint
type MulticastRouterRendezvousPointInventoryView struct {
	Uuid                string    `json:"uuid,omitempty"`
	MulticastRouterUuid string    `json:"multicastRouterUuid,omitempty"`
	RpAddress           string    `json:"rpAddress,omitempty"`
	GroupAddress        string    `json:"groupAddress,omitempty"`
	CreateDate          time.Time `json:"createDate,omitempty"`
	LastOpDate          time.Time `json:"lastOpDate,omitempty"`
}
