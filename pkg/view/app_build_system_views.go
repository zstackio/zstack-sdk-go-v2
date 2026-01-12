// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AppBuildSystemInventoryView AppBuildSystem
type AppBuildSystemInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	StorageType *string `json:"storageType,omitempty"`
	Url *string `json:"url,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Username *string `json:"username,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Status *string `json:"status,omitempty"`
	State *string `json:"state,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
}

// AddAppBuildSystemEventView AddAppBuildSystemEvent
type AddAppBuildSystemEventView struct {
	Inventory AppBuildSystemInventoryView `json:"inventory,omitempty"`
}

// QueryAppBuildSystemView QueryAppBuildSystem
type QueryAppBuildSystemView struct {
	Inventories []AppBuildSystemInventoryView `json:"inventories,omitempty"`
}

// ChangeAppBuildSystemStateEventView ChangeAppBuildSystemStateEvent
type ChangeAppBuildSystemStateEventView struct {
	Inventory AppBuildSystemInventoryView `json:"inventory,omitempty"`
}

// ReconnectAppBuildSystemEventView ReconnectAppBuildSystemEvent
type ReconnectAppBuildSystemEventView struct {
	Inventory AppBuildSystemInventoryView `json:"inventory,omitempty"`
}

// UpdateAppBuildSystemEventView UpdateAppBuildSystemEvent
type UpdateAppBuildSystemEventView struct {
	Inventory AppBuildSystemInventoryView `json:"inventory,omitempty"`
}

// DeleteAppBuildSystemEventView DeleteAppBuildSystemEvent
type DeleteAppBuildSystemEventView struct {
	Success bool `json:"success,omitempty"`
}

