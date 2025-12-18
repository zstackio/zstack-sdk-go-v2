// Copyright (c) ZStack.io, Inc.

package param

// SyncConnectionAccessPointFromRemoteDetailParam SyncConnectionAccessPointFromRemote detail param
type SyncConnectionAccessPointFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	AccessPointId string `json:"accessPointId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncConnectionAccessPointFromRemoteParam SyncConnectionAccessPointFromRemote request param
type SyncConnectionAccessPointFromRemoteParam struct {
	BaseParam
	Params SyncConnectionAccessPointFromRemoteDetailParam `json:"params"`
}
