// Copyright (c) ZStack.io, Inc.

package param

// AddConnectionAccessPointFromRemoteDetailParam AddConnectionAccessPointFromRemote detail param
type AddConnectionAccessPointFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	AccessPointId string `json:"accessPointId" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddConnectionAccessPointFromRemoteParam AddConnectionAccessPointFromRemote request param
type AddConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params AddConnectionAccessPointFromRemoteDetailParam `json:"params"`
}
