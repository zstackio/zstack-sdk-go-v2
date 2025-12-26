// Copyright (c) ZStack.io, Inc.

package param

// UpdateDatasetsDetailParam UpdateDatasets detail param
type UpdateDatasetsDetailParam struct {
	UpdateDatasetStructs []UpdateDatasetStructParam `json:"updateDatasetStructs" validate:"required"`
}

// UpdateDatasetsParam UpdateDatasets request param
type UpdateDatasetsParam struct {
	BaseParam
	Params UpdateDatasetsDetailParam `json:"params"`
}
