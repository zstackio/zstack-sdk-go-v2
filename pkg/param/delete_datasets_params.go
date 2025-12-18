// Copyright (c) ZStack.io, Inc.

package param

// DeleteDatasetsDetailParam DeleteDatasets detail param
type DeleteDatasetsDetailParam struct {
	Uuids []string `json:"uuids" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDatasetsParam DeleteDatasets request param
type DeleteDatasetsParam struct {
	BaseParam
	Params DeleteDatasetsDetailParam `json:"params"`
}
