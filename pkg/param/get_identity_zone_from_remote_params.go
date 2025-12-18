// Copyright (c) ZStack.io, Inc.

package param

// GetIdentityZoneFromRemoteDetailParam GetIdentityZoneFromRemote detail param
type GetIdentityZoneFromRemoteDetailParam struct {
	Type string `json:"type,omitempty"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	RegionId string `json:"regionId,omitempty"`
}

// GetIdentityZoneFromRemoteParam GetIdentityZoneFromRemote request param
type GetIdentityZoneFromRemoteParam struct {
	BaseParam
	Params GetIdentityZoneFromRemoteDetailParam `json:"params"`
}
