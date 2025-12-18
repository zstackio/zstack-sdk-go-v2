// Copyright (c) ZStack.io, Inc.

package param

// GetResourceNamesDetailParam GetResourceNames详细参数
type GetResourceNamesDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
}

// GetResourceNamesParam GetResourceNames请求参数
type GetResourceNamesParam struct {
	BaseParam
	Params GetResourceNamesDetailParam `json:"params"` // 详细参数
}

