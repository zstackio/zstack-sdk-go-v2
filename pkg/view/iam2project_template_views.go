// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectTemplateInventoryView IAM2ProjectTemplate
type IAM2ProjectTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Template TemplateView `json:"template,omitempty"`
}

// CreateIAM2ProjectTemplateEventView CreateIAM2ProjectTemplateEvent
type CreateIAM2ProjectTemplateEventView struct {
	Inventory IAM2ProjectTemplateInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2ProjectTemplateView QueryIAM2ProjectTemplate
type QueryIAM2ProjectTemplateView struct {
	Inventories []IAM2ProjectTemplateInventoryView `json:"inventories,omitempty"`
}

// CreateIAM2ProjectTemplateFromProjectEventView CreateIAM2ProjectTemplateFromProjectEvent
type CreateIAM2ProjectTemplateFromProjectEventView struct {
	Inventory IAM2ProjectTemplateInventoryView `json:"inventory,omitempty"`
}

// UpdateIAM2ProjectTemplateEventView UpdateIAM2ProjectTemplateEvent
type UpdateIAM2ProjectTemplateEventView struct {
	Inventory IAM2ProjectTemplateInventoryView `json:"inventory,omitempty"`
}

// DeleteIAM2ProjectTemplateEventView DeleteIAM2ProjectTemplateEvent
type DeleteIAM2ProjectTemplateEventView struct {
	Success bool `json:"success,omitempty"`
}

