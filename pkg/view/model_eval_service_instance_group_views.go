// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelEvalServiceInstanceGroupInventoryView ModelEvalServiceInstanceGroup
type ModelEvalServiceInstanceGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Limits *int `json:"limits,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`
	TopK *int `json:"topK,omitempty"`
	TopP *float32 `json:"topP,omitempty"`
	MaxLength *int `json:"maxLength,omitempty"`
	MaxNewTokens *int `json:"maxNewTokens,omitempty"`
	RepetitionPenalty *float32 `json:"repetitionPenalty,omitempty"`
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

// QueryModelServiceInstanceGroupView QueryModelServiceInstanceGroup
type QueryModelServiceInstanceGroupView struct {
	Inventories []ModelServiceInstanceGroupInventoryView `json:"inventories,omitempty"`
}

