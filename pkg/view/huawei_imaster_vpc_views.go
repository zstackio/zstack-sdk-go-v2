// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HuaweiIMasterVpcInventoryView HuaweiIMasterVpc
type HuaweiIMasterVpcInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	FabricId string `json:"fabricId,omitempty"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	State string `json:"state,omitempty"`
	IsVpcDeployed bool `json:"isVpcDeployed,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QueryHuaweiIMasterVpcView QueryHuaweiIMasterVpc
type QueryHuaweiIMasterVpcView struct {
	Inventories []HuaweiIMasterVpcInventoryView `json:"inventories,omitempty"`
}

// DeleteHuaweiIMasterVpcEventView DeleteHuaweiIMasterVpcEvent
type DeleteHuaweiIMasterVpcEventView struct {
	Success bool `json:"success,omitempty"`
}

