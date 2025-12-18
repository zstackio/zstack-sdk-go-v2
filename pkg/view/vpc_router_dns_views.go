// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcRouterDnsInventoryView VpcRouterDns
type VpcRouterDnsInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vpcRouterUuid,omitempty"`
	rest string `json:"dns,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

