// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RemoteVtepInventoryView RemoteVtep
type RemoteVtepInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	Port int `json:"port,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
}

// CreateVxlanPoolRemoteVtepEventView CreateVxlanPoolRemoteVtepEvent
type CreateVxlanPoolRemoteVtepEventView struct {
	Inventory RemoteVtepInventoryView `json:"inventory,omitempty"`
}

