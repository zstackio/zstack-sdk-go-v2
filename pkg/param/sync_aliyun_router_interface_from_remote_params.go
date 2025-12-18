// Copyright (c) ZStack.io, Inc.

package param

// SyncAliyunRouterInterfaceFromRemoteDetailParam SyncAliyunRouterInterfaceFromRemote detail param
type SyncAliyunRouterInterfaceFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncAliyunRouterInterfaceFromRemoteParam SyncAliyunRouterInterfaceFromRemote request param
type SyncAliyunRouterInterfaceFromRemoteParam struct {
	BaseParam
	Params SyncAliyunRouterInterfaceFromRemoteDetailParam `json:"params"`
}
