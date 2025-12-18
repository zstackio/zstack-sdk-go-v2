// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunRouterInterfaceFromRemoteDetailParam SyncAliyunRouterInterfaceFromRemote详细参数
type SyncAliyunRouterInterfaceFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouterInterfaceFromRemoteParam SyncAliyunRouterInterfaceFromRemote请求参数
type SyncAliyunRouterInterfaceFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouterInterfaceFromRemoteDetailParam `json:"params"` // 详细参数
}

