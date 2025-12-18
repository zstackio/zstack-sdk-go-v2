// Copyright (c) ZStack.io, Inc.

package param

// GetDataCenterFromRemoteDetailParam GetDataCenterFromRemote详细参数
type GetDataCenterFromRemoteDetailParam struct {
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"endpoint,omitempty"`
}

// GetDataCenterFromRemoteParam GetDataCenterFromRemote请求参数
type GetDataCenterFromRemoteParam struct {
	BaseParam
	Params GetDataCenterFromRemoteDetailParam `json:"params"` // 详细参数
}

