// Copyright (c) ZStack.io, Inc.

package view

// GetObservabilityServerServiceDataView GetObservabilityServerServiceData
type GetObservabilityServerServiceDataView struct {
	Inventories []ObservabilityServerServiceDataInventoryView `json:"inventories,omitempty"`
	Total int `json:"total,omitempty"`
	Success bool `json:"success,omitempty"`
}

