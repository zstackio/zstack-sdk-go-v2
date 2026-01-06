// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcRouterDnsInventoryView VpcRouterDns
type VpcRouterDnsInventoryView struct {
	Id int64 `json:"id,omitempty"`
	VpcRouterUuid string `json:"vpcRouterUuid,omitempty"`
	Dns string `json:"dns,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

