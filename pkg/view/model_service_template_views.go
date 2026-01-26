// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceTemplateInventoryView ModelServiceTemplate
type ModelServiceTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	ModelServiceUuid string `json:"modelServiceUuid,omitempty"`
	CpuArchitecture string `json:"cpuArchitecture,omitempty"`
	VmImageUuid string `json:"vmImageUuid,omitempty"`
	DockerImage string `json:"dockerImage,omitempty"`
}

