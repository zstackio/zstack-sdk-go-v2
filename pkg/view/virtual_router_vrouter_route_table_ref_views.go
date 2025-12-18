// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VirtualRouterVRouterRouteTableRefInventoryView VirtualRouterVRouterRouteTableRef
type VirtualRouterVRouterRouteTableRefInventoryView struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid,omitempty"`
	RouteTableUuid string `json:"routeTableUuid,omitempty"`
}

