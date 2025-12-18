// Copyright (c) ZStack.io, Inc.

package param

// DeleteDatasetsDetailParam DeleteDatasets详细参数
type DeleteDatasetsDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteDatasetsParam DeleteDatasets请求参数
type DeleteDatasetsParam struct {
	BaseParam
	Params DeleteDatasetsDetailParam `json:"params"` // 详细参数
}

