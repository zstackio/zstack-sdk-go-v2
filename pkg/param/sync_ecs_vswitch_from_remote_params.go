// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsVSwitchFromRemoteDetailParam SyncEcsVSwitchFromRemote详细参数
type SyncEcsVSwitchFromRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"vSwitchId,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SyncEcsVSwitchFromRemoteParam SyncEcsVSwitchFromRemote请求参数
type SyncEcsVSwitchFromRemoteParam struct {
	BaseParam
	Params SyncEcsVSwitchFromRemoteDetailParam `json:"params"` // 详细参数
}

