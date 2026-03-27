// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IdentityZoneInventoryView IdentityZone
type IdentityZoneInventoryView struct {
	BaseInfoView
	BaseTimeView
	Closed string `json:"closed,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid,omitempty"`
	ZoneId string `json:"zoneId,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneName string `json:"zoneName,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryIdentityZoneFromLocalView QueryIdentityZoneFromLocal
type QueryIdentityZoneFromLocalView struct {
	Inventories []IdentityZoneInventoryView `json:"inventories,omitempty"`
}

// AddIdentityZoneFromRemoteEventView AddIdentityZoneFromRemoteEvent
type AddIdentityZoneFromRemoteEventView struct {
	Inventory IdentityZoneInventoryView `json:"inventory,omitempty"`
}

