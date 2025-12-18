// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VRouterRouteTableInventoryView VRouterRouteTable
type VRouterRouteTableInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	AttachedRouterRefs []VirtualRouterVRouterRouteTableRefInventoryView `json:"attachedRouterRefs,omitempty"`
	RouteEntries []VRouterRouteEntryInventoryView `json:"routeEntries,omitempty"`
}

