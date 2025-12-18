// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunRouteEntryFromRemoteDetailParam SyncAliyunRouteEntryFromRemote详细参数
type SyncAliyunRouteEntryFromRemoteDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest string `json:"vRouterType" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouteEntryFromRemoteParam SyncAliyunRouteEntryFromRemote请求参数
type SyncAliyunRouteEntryFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouteEntryFromRemoteDetailParam `json:"params"` // 详细参数
}

