// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BuildApplicationInventoryView BuildApplication
type BuildApplicationInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	BuildSystemUuid string `json:"buildSystemUuid,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	AppMetaData string `json:"appMetaData,omitempty"`
	AppId string `json:"appId,omitempty"`
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
}

// CreateBuildAppEventView CreateBuildAppEvent
type CreateBuildAppEventView struct {
	Inventory BuildApplicationInventoryView `json:"inventory,omitempty"`
}

// QueryBuildAppView QueryBuildApp
type QueryBuildAppView struct {
	Inventories []BuildApplicationInventoryView `json:"inventories,omitempty"`
}

// AddBuildAppEventView AddBuildAppEvent
type AddBuildAppEventView struct {
	Inventory BuildApplicationInventoryView `json:"inventory,omitempty"`
}

// UpdateBuildAppEventView UpdateBuildAppEvent
type UpdateBuildAppEventView struct {
	Inventory BuildApplicationInventoryView `json:"inventory,omitempty"`
}

