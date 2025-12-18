// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcHaGroupNetworkServiceRefInventoryView VpcHaGroupNetworkServiceRef
type VpcHaGroupNetworkServiceRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"vpcHaRouterUuid,omitempty"`
	rest string `json:"networkServiceName,omitempty"`
	rest string `json:"networkServiceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

