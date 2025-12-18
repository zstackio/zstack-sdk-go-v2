// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VRouterRouteTableInventoryView VRouterRouteTable
type VRouterRouteTableInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VirtualRouterVRouterRouteTableRefInventoryView `json:"attachedRouterRefs,omitempty"`
	rest []VRouterRouteEntryInventoryView `json:"routeEntries,omitempty"`
}

