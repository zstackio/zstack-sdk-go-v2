// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterTenantFabricRefInventoryView HuaweiIMasterTenantFabricRef
type HuaweiIMasterTenantFabricRefInventoryView struct {
	TenantUuid string `json:"tenantUuid,omitempty"`
	FabricUuid string `json:"fabricUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

