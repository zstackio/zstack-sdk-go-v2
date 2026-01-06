// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AppBuildSystemZoneRefInventoryView AppBuildSystemZoneRef
type AppBuildSystemZoneRefInventoryView struct {
	Id int64 `json:"id,omitempty"`
	BuildSystemUuid string `json:"buildSystemUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// AttachAppBuildSystemToZoneEventView AttachAppBuildSystemToZoneEvent
type AttachAppBuildSystemToZoneEventView struct {
	Inventory AppBuildSystemZoneRefInventoryView `json:"inventory,omitempty"`
}

