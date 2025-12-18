// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AppBuildSystemZoneRefInventoryView AppBuildSystemZoneRef
type AppBuildSystemZoneRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"buildSystemUuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

