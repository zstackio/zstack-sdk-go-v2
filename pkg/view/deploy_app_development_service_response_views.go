// Copyright (c) ZStack.io, Inc.

package view

// DeployAppDevelopmentServiceEventView DeployAppDevelopmentServiceEvent
type DeployAppDevelopmentServiceEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
	App ApplicationDevelopmentServiceInventoryView `json:"app,omitempty"`
	Success bool `json:"success,omitempty"`
}

