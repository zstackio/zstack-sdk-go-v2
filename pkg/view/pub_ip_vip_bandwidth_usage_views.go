// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PubIpVipBandwidthUsageInventoryView PubIpVipBandwidthUsage
type PubIpVipBandwidthUsageInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"accountUuid,omitempty"`
	rest int64 `json:"dateInLong,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"inventory,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"vipName,omitempty"`
	rest string `json:"vipIp,omitempty"`
	rest int64 `json:"bandwidthIn,omitempty"`
	rest int64 `json:"bandwidthOut,omitempty"`
	rest string `json:"vipStatus,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
}

