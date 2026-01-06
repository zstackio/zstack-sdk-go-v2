// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelCenterInventoryView ModelCenter
type ModelCenterInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Status string `json:"status,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	ManagementPort int `json:"managementPort,omitempty"`
	StorageNetworkUuid string `json:"storageNetworkUuid,omitempty"`
	ServiceNetworkUuid string `json:"serviceNetworkUuid,omitempty"`
	ContainerRegistry string `json:"containerRegistry,omitempty"`
	ContainerStorageNetwork string `json:"containerStorageNetwork,omitempty"`
	ContainerNetwork string `json:"containerNetwork,omitempty"`
	Capacity ModelCenterCapacityInventoryView `json:"capacity,omitempty"`
	Zdfs ZdfsInventoryView `json:"zdfs,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// DeleteModelCenterEventView DeleteModelCenterEvent
type DeleteModelCenterEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryModelCenterView QueryModelCenter
type QueryModelCenterView struct {
	Inventories []ModelCenterInventoryView `json:"inventories,omitempty"`
}

// AddModelCenterEventView AddModelCenterEvent
type AddModelCenterEventView struct {
	Inventory ModelCenterInventoryView `json:"inventory,omitempty"`
}

// UpdateModelCenterEventView UpdateModelCenterEvent
type UpdateModelCenterEventView struct {
	Inventory ModelCenterInventoryView `json:"inventory,omitempty"`
}

