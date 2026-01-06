// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteTableVRouterRefInventoryView PolicyRouteTableVRouterRef
type PolicyRouteTableVRouterRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryPolicyRouteTableVRouterRefView QueryPolicyRouteTableVRouterRef
type QueryPolicyRouteTableVRouterRefView struct {
	Inventories []PolicyRouteTableVRouterRefInventoryView `json:"inventories,omitempty"`
}

