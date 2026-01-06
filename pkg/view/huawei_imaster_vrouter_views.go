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
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// CreateHuaweiIMasterVRouterEventView CreateHuaweiIMasterVRouterEvent
type CreateHuaweiIMasterVRouterEventView struct {
	Inventory HuaweiIMasterVRouterInventoryView `json:"inventory,omitempty"`
}

// QueryHuaweiIMasterVRouterView QueryHuaweiIMasterVRouter
type QueryHuaweiIMasterVRouterView struct {
	Inventories []HuaweiIMasterVRouterInventoryView `json:"inventories,omitempty"`
}

// DeleteHuaweiIMasterVRouterEventView DeleteHuaweiIMasterVRouterEvent
type DeleteHuaweiIMasterVRouterEventView struct {
	Success bool `json:"success,omitempty"`
}

