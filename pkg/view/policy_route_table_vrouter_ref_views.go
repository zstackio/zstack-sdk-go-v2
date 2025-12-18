// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteTableVRouterRefInventoryView PolicyRouteTableVRouterRef
type PolicyRouteTableVRouterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

