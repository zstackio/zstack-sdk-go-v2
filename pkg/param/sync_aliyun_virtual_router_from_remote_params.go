// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunVirtualRouterFromRemoteDetailParam SyncAliyunVirtualRouterFromRemote详细参数
type SyncAliyunVirtualRouterFromRemoteDetailParam struct {
	rest string `json:"vpcUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncAliyunVirtualRouterFromRemoteParam SyncAliyunVirtualRouterFromRemote请求参数
type SyncAliyunVirtualRouterFromRemoteParam struct {
	BaseParam
	Params SyncAliyunVirtualRouterFromRemoteDetailParam `json:"params"` // 详细参数
}

