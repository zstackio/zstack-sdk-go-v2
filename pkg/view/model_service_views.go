// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelServiceInventoryView ModelService
type ModelServiceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"readme,omitempty"`
	rest string `json:"yaml,omitempty"`
	rest int `json:"requestCpu,omitempty"`
	rest int64 `json:"requestMemory,omitempty"`
	rest string `json:"modelCenterUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"framework,omitempty"`
	rest string `json:"source,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest bool `json:"system,omitempty"`
	rest string `json:"gpuComputeCapability,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"pythonVersion,omitempty"`
	rest string `json:"condaVersion,omitempty"`
	rest string `json:"startCommand,omitempty"`
	rest string `json:"containerCommand,omitempty"`
	rest string `json:"containerArgs,omitempty"`
	rest bool `json:"supportDistributed,omitempty"`
	rest []string `json:"cpuArchitectures,omitempty"`
	rest map[string]interface{} `json:"vendorToSpecUuidsMap,omitempty"`
	rest []ModelServiceRefInventoryView `json:"modelServiceRefs,omitempty"`
	rest []ModelServiceTemplateInventoryView `json:"modelServiceImages,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

