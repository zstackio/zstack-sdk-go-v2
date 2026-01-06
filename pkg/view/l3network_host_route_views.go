// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L3NetworkHostRouteInventoryView L3NetworkHostRoute
type L3NetworkHostRouteInventoryView struct {
	Id int64 `json:"id,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Nexthop string `json:"nexthop,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

