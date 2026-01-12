// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HostNetworkLabelInventoryView HostNetworkLabel
type HostNetworkLabelInventoryView struct {
	BaseInfoView
	BaseTimeView
	ServiceType *string `json:"serviceType,omitempty"`
	System *bool `json:"system,omitempty"`
}

// CreateHostNetworkServiceTypeEventView CreateHostNetworkServiceTypeEvent
type CreateHostNetworkServiceTypeEventView struct {
	Inventory HostNetworkLabelInventoryView `json:"inventory,omitempty"`
}

// UpdateHostNetworkServiceTypeEventView UpdateHostNetworkServiceTypeEvent
type UpdateHostNetworkServiceTypeEventView struct {
	Inventory HostNetworkLabelInventoryView `json:"inventory,omitempty"`
}

