// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelServiceTemplateInventoryView ModelServiceTemplate
type ModelServiceTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"modelServiceUuid,omitempty"`
	rest string `json:"cpuArchitecture,omitempty"`
	rest string `json:"vmImageUuid,omitempty"`
	rest string `json:"dockerImage,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

