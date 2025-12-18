// Copyright (c) ZStack.io, Inc.

package param

// SyncVirtualBorderRouterFromRemoteDetailParam SyncVirtualBorderRouterFromRemote detail param
type SyncVirtualBorderRouterFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncVirtualBorderRouterFromRemoteParam SyncVirtualBorderRouterFromRemote request param
type SyncVirtualBorderRouterFromRemoteParam struct {
	BaseParam
	Params SyncVirtualBorderRouterFromRemoteDetailParam `json:"params"`
}
