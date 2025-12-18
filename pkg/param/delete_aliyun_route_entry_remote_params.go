// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunRouteEntryRemoteDetailParam DeleteAliyunRouteEntryRemote detail param
type DeleteAliyunRouteEntryRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunRouteEntryRemoteParam DeleteAliyunRouteEntryRemote request param
type DeleteAliyunRouteEntryRemoteParam struct {
	BaseParam
	Params DeleteAliyunRouteEntryRemoteDetailParam `json:"params"`
}
