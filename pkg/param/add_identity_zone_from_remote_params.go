// Copyright (c) ZStack.io, Inc.

package param

// AddIdentityZoneFromRemoteDetailParam AddIdentityZoneFromRemote详细参数
type AddIdentityZoneFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"zoneId,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddIdentityZoneFromRemoteParam AddIdentityZoneFromRemote请求参数
type AddIdentityZoneFromRemoteParam struct {
	BaseParam
	Params AddIdentityZoneFromRemoteDetailParam `json:"params"` // 详细参数
}

