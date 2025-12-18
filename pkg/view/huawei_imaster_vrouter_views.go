// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterVRouterInventoryView HuaweiIMasterVRouter
type HuaweiIMasterVRouterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LogicalNetworkId string `json:"logicalNetworkId,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	FabricUuid string `json:"fabricUuid,omitempty"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

