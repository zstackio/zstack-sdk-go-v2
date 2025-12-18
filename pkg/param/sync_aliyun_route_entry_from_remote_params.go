// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunRouteEntryFromRemoteDetailParam SyncAliyunRouteEntryFromRemote detail param
type SyncAliyunRouteEntryFromRemoteDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouteEntryFromRemoteParam SyncAliyunRouteEntryFromRemote request param
type SyncAliyunRouteEntryFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouteEntryFromRemoteDetailParam `json:"params"`
}
