// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZdfsInventoryView Zdfs
type ZdfsInventoryView struct {
	BaseInfoView
	BaseTimeView
	ZoneUuid string                   `json:"zoneUuid,omitempty"`
	Url      string                   `json:"url,omitempty"`
	Status   string                   `json:"status,omitempty"`
	HostName string                   `json:"hostName,omitempty"`
	SshPort  int                      `json:"sshPort,omitempty"`
	Storage  ZdfsStorageInventoryView `json:"storage,omitempty"`
}

// ReconnectZdfsEventView ReconnectZdfsEvent
type ReconnectZdfsEventView struct {
	Inventory ZdfsInventoryView `json:"inventory,omitempty"`
}

// QueryZdfsView QueryZdfs
type QueryZdfsView struct {
	Inventories []ZdfsInventoryView `json:"inventories,omitempty"`
}
