// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelsDetailParam DeleteModels详细参数
type DeleteModelsDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteModelsParam DeleteModels请求参数
type DeleteModelsParam struct {
	BaseParam
	Params DeleteModelsDetailParam `json:"params"` // 详细参数
}

