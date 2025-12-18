// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourceConfigsDetailParam UpdateResourceConfigs详细参数
type UpdateResourceConfigsDetailParam struct {
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest []interface{} `json:"resourceConfigs" validate:"required"` // 必填
}

// UpdateResourceConfigsParam UpdateResourceConfigs请求参数
type UpdateResourceConfigsParam struct {
	BaseParam
	Params UpdateResourceConfigsDetailParam `json:"params"` // 详细参数
}

