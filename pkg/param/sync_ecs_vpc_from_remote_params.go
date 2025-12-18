// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsVpcFromRemoteDetailParam SyncEcsVpcFromRemote detail param
type SyncEcsVpcFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	EcsVpcId string `json:"ecsVpcId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsVpcFromRemoteParam SyncEcsVpcFromRemote request param
type SyncEcsVpcFromRemoteParam struct {
	BaseParam
	Params SyncEcsVpcFromRemoteDetailParam `json:"params"`
}
