// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ExternalServiceConfigurationInventoryView ExternalServiceConfiguration
type ExternalServiceConfigurationInventoryView struct {
	BaseInfoView
	BaseTimeView
	ServiceType string `json:"serviceType,omitempty"`
	Configuration string `json:"configuration,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryExternalServiceConfigurationView QueryExternalServiceConfiguration
type QueryExternalServiceConfigurationView struct {
	Inventories []ExternalServiceConfigurationInventoryView `json:"inventories,omitempty"`
}

// UpdateExternalServiceConfigurationEventView UpdateExternalServiceConfigurationEvent
type UpdateExternalServiceConfigurationEventView struct {
	Inventory ExternalServiceConfigurationInventoryView `json:"inventory,omitempty"`
}

// AddExternalServiceConfigurationEventView AddExternalServiceConfigurationEvent
type AddExternalServiceConfigurationEventView struct {
	Inventory ExternalServiceConfigurationInventoryView `json:"inventory,omitempty"`
}

// DeleteExternalServiceConfigurationEventView DeleteExternalServiceConfigurationEvent
type DeleteExternalServiceConfigurationEventView struct {
	Success bool `json:"success,omitempty"`
}

