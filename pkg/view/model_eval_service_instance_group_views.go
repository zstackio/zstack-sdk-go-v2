// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ModelEvalServiceInstanceGroupInventoryView ModelEvalServiceInstanceGroup
type ModelEvalServiceInstanceGroupInventoryView struct {
	rest int `json:"limits,omitempty"`
	rest float32 `json:"temperature,omitempty"`
	rest int `json:"topK,omitempty"`
	rest float32 `json:"topP,omitempty"`
	rest int `json:"maxLength,omitempty"`
	rest int `json:"maxNewTokens,omitempty"`
	rest float32 `json:"repetitionPenalty,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"modelServiceUuid,omitempty"`
	rest string `json:"modelUuid,omitempty"`
	rest []ModelServiceInstanceInventoryView `json:"instances,omitempty"`
	rest []ModelServiceGroupDatasetRefInventoryView `json:"datasetRefInventories,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"modelServiceType,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"yaml,omitempty"`
	rest []string `json:"supportMetrics,omitempty"`
}

