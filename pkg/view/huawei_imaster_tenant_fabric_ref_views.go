// Copyright (c) ZStack.io, Inc.

package view

import "time"

// HuaweiIMasterTenantFabricRefInventoryView HuaweiIMasterTenantFabricRef
type HuaweiIMasterTenantFabricRefInventoryView struct {
	rest string `json:"tenantUuid,omitempty"`
	rest string `json:"fabricUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

