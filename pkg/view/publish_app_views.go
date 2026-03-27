// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PublishAppInventoryView PublishApp
type PublishAppInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	AppMetaData string `json:"appMetaData,omitempty"`
	PreParams string `json:"preParams,omitempty"`
	VmRelationShip string `json:"vmRelationShip,omitempty"`
	BuildAppUuid string `json:"buildAppUuid,omitempty"`
	Type string `json:"type,omitempty"`
	AppId string `json:"appId,omitempty"`
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
}

// QueryPublishAppView QueryPublishApp
type QueryPublishAppView struct {
	Inventories []PublishAppInventoryView `json:"inventories,omitempty"`
}

// UpdatePublishAppEventView UpdatePublishAppEvent
type UpdatePublishAppEventView struct {
	Inventory PublishAppInventoryView `json:"inventory,omitempty"`
}

// DeletePublishAppEventView DeletePublishAppEvent
type DeletePublishAppEventView struct {
	Success bool `json:"success,omitempty"`
}

