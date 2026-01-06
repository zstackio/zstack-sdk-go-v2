// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZdfsInventoryView Zdfs
type ZdfsInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Status string `json:"status,omitempty"`
	HostName string `json:"hostName,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Storage ZdfsStorageInventoryView `json:"storage,omitempty"`
}

// ReconnectZdfsEventView ReconnectZdfsEvent
type ReconnectZdfsEventView struct {
	Inventory ZdfsInventoryView `json:"inventory,omitempty"`
}

// QueryZdfsView QueryZdfs
type QueryZdfsView struct {
	Inventories []ZdfsInventoryView `json:"inventories,omitempty"`
}

