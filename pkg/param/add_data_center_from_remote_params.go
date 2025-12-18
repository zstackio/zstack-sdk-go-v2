// Copyright (c) ZStack.io, Inc.

package param

// AddDataCenterFromRemoteDetailParam AddDataCenterFromRemote detail param
type AddDataCenterFromRemoteDetailParam struct {
	RegionId string `json:"regionId" validate:"required"`
	Type string `json:"type" validate:"required"`
	SyncZones bool `json:"syncZones,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddDataCenterFromRemoteParam AddDataCenterFromRemote request param
type AddDataCenterFromRemoteParam struct {
	BaseParam
	Params AddDataCenterFromRemoteDetailParam `json:"params"`
}
