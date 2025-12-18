// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HuaweiIMasterVRouterInventoryView HuaweiIMasterVRouter
type HuaweiIMasterVRouterInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"logicalNetworkId,omitempty"`
	rest string `json:"tenantId,omitempty"`
	rest string `json:"fabricUuid,omitempty"`
	rest string `json:"sdnControllerUuid,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

