// Copyright (c) ZStack.io, Inc.

package param

// UpdateDatasetDetailParam UpdateDataset详细参数
type UpdateDatasetDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest []string `json:"usageScenarios,omitempty"`
	rest string `json:"dataType,omitempty"`
}

// UpdateDatasetParam UpdateDataset请求参数
type UpdateDatasetParam struct {
	BaseParam
	Params UpdateDatasetDetailParam `json:"params"` // 详细参数
}

