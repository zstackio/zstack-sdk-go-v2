// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ApplicationDevelopmentServiceInventoryView ApplicationDevelopmentService
type ApplicationDevelopmentServiceInventoryView struct {
	BaseInfoView
	BaseTimeView
	DeploymentStatus *string `json:"deploymentStatus,omitempty"`
	Service ModelServiceInventoryView `json:"service,omitempty"`
	ModelServiceUuid *string `json:"modelServiceUuid,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	Instances []ModelServiceInstanceInventoryView `json:"instances,omitempty"`
	DatasetRefInventories []ModelServiceGroupDatasetRefInventoryView `json:"datasetRefInventories,omitempty"`
	Status *string `json:"status,omitempty"`
	ModelServiceType *string `json:"modelServiceType,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Yaml *string `json:"yaml,omitempty"`
	SupportMetrics []string `json:"supportMetrics,omitempty"`
}

// QueryApplicationDevelopmentServiceView QueryApplicationDevelopmentService
type QueryApplicationDevelopmentServiceView struct {
	Inventories []ApplicationDevelopmentServiceInventoryView `json:"inventories,omitempty"`
}

