// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunRouteEntryRemoteDetailParam DeleteAliyunRouteEntryRemote详细参数
type DeleteAliyunRouteEntryRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouteEntryRemoteParam DeleteAliyunRouteEntryRemote请求参数
type DeleteAliyunRouteEntryRemoteParam struct {
	BaseParam
	Params DeleteAliyunRouteEntryRemoteDetailParam `json:"params"` // 详细参数
}

