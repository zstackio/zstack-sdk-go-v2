// Copyright (c) ZStack.io, Inc.

package param

// CheckScsiLunClusterStatusDetailParam CheckScsiLunClusterStatus详细参数
type CheckScsiLunClusterStatusDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// CheckScsiLunClusterStatusParam CheckScsiLunClusterStatus请求参数
type CheckScsiLunClusterStatusParam struct {
	BaseParam
	Params CheckScsiLunClusterStatusDetailParam `json:"params"` // 详细参数
}

