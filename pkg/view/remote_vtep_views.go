// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RemoteVtepInventoryView RemoteVtep
type RemoteVtepInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClusterUuid string `json:"clusterUuid,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	Port int `json:"port,omitempty"`
	Type string `json:"type,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
}

// CreateVxlanPoolRemoteVtepEventView CreateVxlanPoolRemoteVtepEvent
type CreateVxlanPoolRemoteVtepEventView struct {
	Inventory RemoteVtepInventoryView `json:"inventory,omitempty"`
}

