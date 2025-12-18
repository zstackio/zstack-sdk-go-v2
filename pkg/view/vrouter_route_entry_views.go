// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VRouterRouteEntryInventoryView VRouterRouteEntry
type VRouterRouteEntryInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	RouteTableUuid string `json:"routeTableUuid,omitempty"`
	Destination string `json:"destination,omitempty"`
	Target string `json:"target,omitempty"`
	Distance int `json:"distance,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

