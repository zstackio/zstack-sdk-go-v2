// Copyright (c) ZStack.io, Inc.

package param

// SyncVirtualBorderRouterFromRemoteDetailParam SyncVirtualBorderRouterFromRemote详细参数
type SyncVirtualBorderRouterFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncVirtualBorderRouterFromRemoteParam SyncVirtualBorderRouterFromRemote请求参数
type SyncVirtualBorderRouterFromRemoteParam struct {
	BaseParam
	Params SyncVirtualBorderRouterFromRemoteDetailParam `json:"params"` // 详细参数
}

