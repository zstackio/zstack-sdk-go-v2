// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ModelServiceInstanceGroupInventoryView ModelServiceInstanceGroup
type ModelServiceInstanceGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelServiceUuid string `json:"modelServiceUuid,omitempty"`
	ModelUuid string `json:"modelUuid,omitempty"`
	Instances []ModelServiceInstanceInventoryView `json:"instances,omitempty"`
	DatasetRefInventories []ModelServiceGroupDatasetRefInventoryView `json:"datasetRefInventories,omitempty"`
	Status string `json:"status,omitempty"`
	ModelServiceType string `json:"modelServiceType,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Yaml string `json:"yaml,omitempty"`
	SupportMetrics []string `json:"supportMetrics,omitempty"`
	ExportPath string `json:"exportPath,omitempty"`
}

// UpdateModelServiceInstanceGroupEventView UpdateModelServiceInstanceGroupEvent
type UpdateModelServiceInstanceGroupEventView struct {
	Inventory ModelServiceInstanceGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteModelServiceInstanceGroupEventView DeleteModelServiceInstanceGroupEvent
type DeleteModelServiceInstanceGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

