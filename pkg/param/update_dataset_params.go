// Copyright (c) ZStack.io, Inc.

package param

// UpdateDatasetDetailParam UpdateDataset detail param
type UpdateDatasetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UsageScenarios []string `json:"usageScenarios,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

// UpdateDatasetParam UpdateDataset request param
type UpdateDatasetParam struct {
	BaseParam
	Params UpdateDatasetDetailParam `json:"params"`
}
