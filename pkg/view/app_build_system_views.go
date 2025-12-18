// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AppBuildSystemInventoryView AppBuildSystem
type AppBuildSystemInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	Url string `json:"url,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Username string `json:"username,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

