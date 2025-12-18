// Copyright (c) ZStack.io, Inc.

package param

// AddIdentityZoneFromRemoteDetailParam AddIdentityZoneFromRemote detail param
type AddIdentityZoneFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	ZoneId string `json:"zoneId,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddIdentityZoneFromRemoteParam AddIdentityZoneFromRemote request param
type AddIdentityZoneFromRemoteParam struct {
	BaseParam
	Params AddIdentityZoneFromRemoteDetailParam `json:"params"`
}
