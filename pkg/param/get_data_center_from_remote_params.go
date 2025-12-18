// Copyright (c) ZStack.io, Inc.

package param

// GetDataCenterFromRemoteDetailParam GetDataCenterFromRemote detail param
type GetDataCenterFromRemoteDetailParam struct {
	Type string `json:"type" validate:"required"`
	Endpoint string `json:"endpoint,omitempty"`
}

// GetDataCenterFromRemoteParam GetDataCenterFromRemote request param
type GetDataCenterFromRemoteParam struct {
	BaseParam
	Params GetDataCenterFromRemoteDetailParam `json:"params"`
}
