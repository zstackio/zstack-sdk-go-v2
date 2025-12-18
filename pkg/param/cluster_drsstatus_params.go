// Copyright (c) ZStack.io, Inc.

package param

// GetClusterDRSStatusDetailParam GetClusterDRSStatus详细参数
type GetClusterDRSStatusDetailParam struct {
	rest string `json:"drsUuid" validate:"required"` // 必填
}

// GetClusterDRSStatusParam GetClusterDRSStatus请求参数
type GetClusterDRSStatusParam struct {
	BaseParam
	Params GetClusterDRSStatusDetailParam `json:"params"` // 详细参数
}

