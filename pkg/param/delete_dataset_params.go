// Copyright (c) ZStack.io, Inc.

package param

// DeleteDatasetDetailParam DeleteDataset detail param
type DeleteDatasetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteDatasetParam DeleteDataset request param
type DeleteDatasetParam struct {
	BaseParam
	Params DeleteDatasetDetailParam `json:"params"`
}
