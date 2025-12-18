// Copyright (c) ZStack.io, Inc.

package param

// DeleteDataCenterInLocalDetailParam DeleteDataCenterInLocal详细参数
type DeleteDataCenterInLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteDataCenterInLocalParam DeleteDataCenterInLocal请求参数
type DeleteDataCenterInLocalParam struct {
	BaseParam
	Params DeleteDataCenterInLocalDetailParam `json:"params"` // 详细参数
}

