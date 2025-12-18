// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunVirtualRouterFromRemoteDetailParam SyncAliyunVirtualRouterFromRemote detail param
type SyncAliyunVirtualRouterFromRemoteDetailParam struct {
	VpcUuid string `json:"vpcUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunVirtualRouterFromRemoteParam SyncAliyunVirtualRouterFromRemote request param
type SyncAliyunVirtualRouterFromRemoteParam struct {
	BaseParam
	Params SyncAliyunVirtualRouterFromRemoteDetailParam `json:"params"`
}
