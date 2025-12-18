// Copyright (c) ZStack.io, Inc.

package param

// CreateDatasetDetailParam CreateDataset detail param
type CreateDatasetDetailParam struct {
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
	Params CreateDatasetDetailParam `json:"params"`
}
