// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteTableVRouterRefInventoryView PolicyRouteTableVRouterRef
type PolicyRouteTableVRouterRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"tableUuid,omitempty"`
	rest string `json:"vRouterUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

