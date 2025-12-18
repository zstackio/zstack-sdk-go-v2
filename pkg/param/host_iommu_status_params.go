// Copyright (c) ZStack.io, Inc.

package param

// GetHostIommuStatusDetailParam GetHostIommuStatus详细参数
type GetHostIommuStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetHostIommuStatusParam GetHostIommuStatus请求参数
type GetHostIommuStatusParam struct {
	BaseParam
	Params GetHostIommuStatusDetailParam `json:"params"` // 详细参数
}

