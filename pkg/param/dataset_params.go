// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateDatasetParamDetail CreateDataset detail param
type CreateDatasetParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	ModelCenterUuid string `json:"modelCenterUuid" validate:"required"`
	Token string `json:"token,omitempty"`
	System bool `json:"system,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDatasetParam CreateDataset request param
type CreateDatasetParam struct {
	BaseParam
	CreateDataset CreateDatasetParamDetail `json:"createDataset"`
}
// DeleteDatasetParamDetail DeleteDataset detail param
type DeleteDatasetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteDatasetParam DeleteDataset request param
type DeleteDatasetParam struct {
	BaseParam
	DeleteDataset DeleteDatasetParamDetail `json:"deleteDataset"`
}
// UpdateDatasetParamDetail UpdateDataset detail param
type UpdateDatasetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UsageScenarios []string `json:"usageScenarios,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

// UpdateDatasetParam UpdateDataset request param
type UpdateDatasetParam struct {
	BaseParam
	UpdateDataset UpdateDatasetParamDetail `json:"updateDataset"`
}
