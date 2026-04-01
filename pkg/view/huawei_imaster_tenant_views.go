// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterTenantInventoryView HuaweiIMasterTenant
type HuaweiIMasterTenantInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	FabricIds []string `json:"fabricIds,omitempty"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	State string `json:"state,omitempty"`
}

// DeleteHuaweiIMasterTenantEventView DeleteHuaweiIMasterTenantEvent
type DeleteHuaweiIMasterTenantEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryHuaweiIMasterTenantView QueryHuaweiIMasterTenant
type QueryHuaweiIMasterTenantView struct {
	Inventories []HuaweiIMasterTenantInventoryView `json:"inventories,omitempty"`
}

