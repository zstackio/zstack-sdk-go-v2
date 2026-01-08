// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcRouterDnsInventoryView VpcRouterDns
type VpcRouterDnsInventoryView struct {
	Id            int64     `json:"id,omitempty"`
	VpcRouterUuid string    `json:"vpcRouterUuid,omitempty"`
	Dns           string    `json:"dns,omitempty"`
	CreateDate    time.Time `json:"createDate,omitempty"`
	LastOpDate    time.Time `json:"lastOpDate,omitempty"`
}
