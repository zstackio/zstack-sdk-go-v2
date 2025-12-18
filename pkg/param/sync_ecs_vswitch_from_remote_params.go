// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsVSwitchFromRemoteDetailParam SyncEcsVSwitchFromRemote detail param
type SyncEcsVSwitchFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	VSwitchId string `json:"vSwitchId,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsVSwitchFromRemoteParam SyncEcsVSwitchFromRemote request param
type SyncEcsVSwitchFromRemoteParam struct {
	BaseParam
	Params SyncEcsVSwitchFromRemoteDetailParam `json:"params"`
}
