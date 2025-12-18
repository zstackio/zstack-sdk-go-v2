// Copyright (c) ZStack.io, Inc.

package param

// GetConnectionAccessPointFromRemoteDetailParam GetConnectionAccessPointFromRemote详细参数
type GetConnectionAccessPointFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
}

// GetConnectionAccessPointFromRemoteParam GetConnectionAccessPointFromRemote请求参数
type GetConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params GetConnectionAccessPointFromRemoteDetailParam `json:"params"` // 详细参数
}

