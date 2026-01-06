// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ContainerManagementEndpointInventoryView ContainerManagementEndpoint
type ContainerManagementEndpointInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AccessKeyId string `json:"accessKeyId,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	ManagementPort int `json:"managementPort,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// UpdateContainerManagementEndpointEventView UpdateContainerManagementEndpointEvent
type UpdateContainerManagementEndpointEventView struct {
	Inventory ContainerManagementEndpointInventoryView `json:"inventory,omitempty"`
}

// QueryContainerManagementEndpointView QueryContainerManagementEndpoint
type QueryContainerManagementEndpointView struct {
	Inventories []ContainerManagementEndpointInventoryView `json:"inventories,omitempty"`
}

// AddContainerManagementEndpointEventView AddContainerManagementEndpointEvent
type AddContainerManagementEndpointEventView struct {
	Inventory ContainerManagementEndpointInventoryView `json:"inventory,omitempty"`
}

// SyncContainerManagementEndpointEventView SyncContainerManagementEndpointEvent
type SyncContainerManagementEndpointEventView struct {
	Inventory ContainerManagementEndpointInventoryView `json:"inventory,omitempty"`
}

// DeleteContainerManagementEndpointEventView DeleteContainerManagementEndpointEvent
type DeleteContainerManagementEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

