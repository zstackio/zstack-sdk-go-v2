// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NvmeServerInventoryView NvmeServer
type NvmeServerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Port *int `json:"port,omitempty"`
	State *string `json:"state,omitempty"`
	Transport *string `json:"transport,omitempty"`
	NvmeTargets []NvmeTargetInventoryView `json:"nvmeTargets,omitempty"`
	NvmeClusterRefs []NvmeServerClusterRefInventoryView `json:"nvmeClusterRefs,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// QueryNvmeServerView QueryNvmeServer
type QueryNvmeServerView struct {
	Inventories []NvmeServerInventoryView `json:"inventories,omitempty"`
}

// DeleteNvmeServerEventView DeleteNvmeServerEvent
type DeleteNvmeServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// AttachNvmeServerToClusterEventView AttachNvmeServerToClusterEvent
type AttachNvmeServerToClusterEventView struct {
	Inventory NvmeServerInventoryView `json:"inventory,omitempty"`
}

// AddNvmeServerEventView AddNvmeServerEvent
type AddNvmeServerEventView struct {
	Inventory NvmeServerInventoryView `json:"inventory,omitempty"`
}

// DetachNvmeServerFromClusterEventView DetachNvmeServerFromClusterEvent
type DetachNvmeServerFromClusterEventView struct {
	Inventory NvmeServerInventoryView `json:"inventory,omitempty"`
}

