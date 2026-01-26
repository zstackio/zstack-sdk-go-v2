// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterTenantFabricRefInventoryView HuaweiIMasterTenantFabricRef
type HuaweiIMasterTenantFabricRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	TenantUuid string `json:"tenantUuid,omitempty"`
	FabricUuid string `json:"fabricUuid,omitempty"`
}

