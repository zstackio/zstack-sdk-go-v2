// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PubIpVmNicBandwidthUsageInventoryView PubIpVmNicBandwidthUsage
type PubIpVmNicBandwidthUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int64 `json:"bandwidthOut,omitempty"`
	rest int64 `json:"bandwidthIn,omitempty"`
	rest string `json:"vmNicIp,omitempty"`
	rest string `json:"vmNicStatus,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
}

