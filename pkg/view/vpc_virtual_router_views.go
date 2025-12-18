// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVirtualRouterInventoryView VpcVirtualRouter
type VpcVirtualRouterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vrId,omitempty"`
	rest string `json:"vpcUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

