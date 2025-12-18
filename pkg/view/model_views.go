// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelInventoryView Model
type ModelInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"parameters,omitempty"`
	rest string `json:"modelCenterUuid,omitempty"`
	rest string `json:"logo,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest string `json:"modelId,omitempty"`
	rest string `json:"introduction,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
	rest int64 `json:"minGpuMemory,omitempty"`
	rest []int `json:"recommendedGpuNum,omitempty"`
	rest string `json:"gpuConstraintDescription,omitempty"`
	rest string `json:"versionSemver,omitempty"`
	rest bool `json:"isLatestVersion,omitempty"`
	rest string `json:"artifactChecksum,omitempty"`
	rest int64 `json:"artifactSizeBytes,omitempty"`
	rest string `json:"architectureType,omitempty"`
	rest string `json:"frameworkVersion,omitempty"`
	rest string `json:"requiredAccelerator,omitempty"`
	rest []ModelServiceRefInventoryView `json:"modelServiceRefs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

