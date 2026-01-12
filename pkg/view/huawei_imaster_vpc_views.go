// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterVpcInventoryView HuaweiIMasterVpc
type HuaweiIMasterVpcInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	FabricId *string `json:"fabricId,omitempty"`
	SdnControllerUuid *string `json:"sdnControllerUuid,omitempty"`
	State *string `json:"state,omitempty"`
	IsVpcDeployed *bool `json:"isVpcDeployed,omitempty"`
}

// QueryHuaweiIMasterVpcView QueryHuaweiIMasterVpc
type QueryHuaweiIMasterVpcView struct {
	Inventories []HuaweiIMasterVpcInventoryView `json:"inventories,omitempty"`
}

// DeleteHuaweiIMasterVpcEventView DeleteHuaweiIMasterVpcEvent
type DeleteHuaweiIMasterVpcEventView struct {
	Success bool `json:"success,omitempty"`
}

