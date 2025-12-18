// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VRouterRouteEntryInventoryView VRouterRouteEntry
type VRouterRouteEntryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"routeTableUuid,omitempty"`
	rest string `json:"destination,omitempty"`
	rest string `json:"target,omitempty"`
	rest int `json:"distance,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

