// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IpUseInventoryView IpUse
type IpUseInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"usedIpUuid,omitempty"`
	rest string `json:"serviceId,omitempty"`
	rest string `json:"use,omitempty"`
	rest string `json:"details,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

