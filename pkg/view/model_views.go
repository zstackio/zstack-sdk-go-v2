// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ModelInventoryView Model
type ModelInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	ModelCenterUuid string `json:"modelCenterUuid,omitempty"`
	Logo string `json:"logo,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	ModelId string `json:"modelId,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Size int64 `json:"size,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	MinGpuMemory int64 `json:"minGpuMemory,omitempty"`
	RecommendedGpuNum []int `json:"recommendedGpuNum,omitempty"`
	GpuConstraintDescription string `json:"gpuConstraintDescription,omitempty"`
	VersionSemver string `json:"versionSemver,omitempty"`
	IsLatestVersion bool `json:"isLatestVersion,omitempty"`
	ArtifactChecksum string `json:"artifactChecksum,omitempty"`
	ArtifactSizeBytes int64 `json:"artifactSizeBytes,omitempty"`
	ArchitectureType string `json:"architectureType,omitempty"`
	FrameworkVersion string `json:"frameworkVersion,omitempty"`
	RequiredAccelerator string `json:"requiredAccelerator,omitempty"`
	ModelServiceRefs []ModelServiceRefInventoryView `json:"modelServiceRefs,omitempty"`
}

// QueryModelView QueryModel
type QueryModelView struct {
	Inventories []ModelInventoryView `json:"inventories,omitempty"`
}

// DeleteModelEventView DeleteModelEvent
type DeleteModelEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateModelEventView UpdateModelEvent
type UpdateModelEventView struct {
	Inventory ModelInventoryView `json:"inventory,omitempty"`
}

// AddModelEventView AddModelEvent
type AddModelEventView struct {
	Inventory ModelInventoryView `json:"inventory,omitempty"`
}

