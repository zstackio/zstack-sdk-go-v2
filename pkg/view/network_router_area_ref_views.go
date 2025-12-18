// Copyright (c) ZStack.io, Inc.

package view

import "time"

// NetworkRouterAreaRefInventoryView NetworkRouterAreaRef
type NetworkRouterAreaRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest string `json:"applianceVmType,omitempty"`
	rest string `json:"routerAreaUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

