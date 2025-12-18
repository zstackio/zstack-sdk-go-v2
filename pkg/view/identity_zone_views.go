// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IdentityZoneInventoryView IdentityZone
type IdentityZoneInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"closed,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"zoneId,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"zoneName,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

