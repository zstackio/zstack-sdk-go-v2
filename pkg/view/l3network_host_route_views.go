// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L3NetworkHostRouteInventoryView L3NetworkHostRoute
type L3NetworkHostRouteInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"prefix,omitempty"`
	rest string `json:"nexthop,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

