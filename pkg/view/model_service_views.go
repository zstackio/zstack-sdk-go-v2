// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceInventoryView ModelService
type ModelServiceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Readme string `json:"readme,omitempty"`
	Yaml string `json:"yaml,omitempty"`
	RequestCpu int `json:"requestCpu,omitempty"`
	RequestMemory int64 `json:"requestMemory,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Framework string `json:"framework,omitempty"`
	Source string `json:"source,omitempty"`
	Size int64 `json:"size,omitempty"`
	System bool `json:"system,omitempty"`
	GpuComputeCapability string `json:"gpuComputeCapability,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	PythonVersion string `json:"pythonVersion,omitempty"`
	CondaVersion string `json:"condaVersion,omitempty"`
	StartCommand string `json:"startCommand,omitempty"`
	ContainerCommand string `json:"containerCommand,omitempty"`
	ContainerArgs string `json:"containerArgs,omitempty"`
	SupportDistributed bool `json:"supportDistributed,omitempty"`
	CpuArchitectures []string `json:"cpuArchitectures,omitempty"`
	VendorToSpecUuidsMap map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	ModelServiceRefs []ModelServiceRefInventoryView `json:"modelServiceRefs,omitempty"`
	ModelServiceImages []ModelServiceTemplateInventoryView `json:"modelServiceImages,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// CloneModelServiceEventView CloneModelServiceEvent
type CloneModelServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// UpdateModelServiceEventView UpdateModelServiceEvent
type UpdateModelServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// DeleteModelServiceEventView DeleteModelServiceEvent
type DeleteModelServiceEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddModelServiceEventView AddModelServiceEvent
type AddModelServiceEventView struct {
	Inventory ModelServiceInventoryView `json:"inventory,omitempty"`
}

// QueryModelServiceView QueryModelService
type QueryModelServiceView struct {
	Inventories []ModelServiceInventoryView `json:"inventories,omitempty"`
}

