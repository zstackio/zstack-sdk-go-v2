// Copyright (c) ZStack.io, Inc.

package param

// GetConnectionAccessPointFromRemoteDetailParam GetConnectionAccessPointFromRemote detail param
type GetConnectionAccessPointFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
}

// GetConnectionAccessPointFromRemoteParam GetConnectionAccessPointFromRemote request param
type GetConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params GetConnectionAccessPointFromRemoteDetailParam `json:"params"`
}
