// Copyright (c) ZStack.io, Inc.

package param

// GetIdentityZoneFromRemoteDetailParam GetIdentityZoneFromRemote详细参数
type GetIdentityZoneFromRemoteDetailParam struct {
	rest string `json:"type,omitempty"`
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"regionId,omitempty"`
}

// GetIdentityZoneFromRemoteParam GetIdentityZoneFromRemote请求参数
type GetIdentityZoneFromRemoteParam struct {
	BaseParam
	Params GetIdentityZoneFromRemoteDetailParam `json:"params"` // 详细参数
}

