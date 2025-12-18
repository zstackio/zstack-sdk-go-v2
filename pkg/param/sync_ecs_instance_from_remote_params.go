// Copyright (c) ZStack.io, Inc.

package param

// SyncEcsInstanceFromRemoteDetailParam SyncEcsInstanceFromRemote detail param
type SyncEcsInstanceFromRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	OnlyZstack bool `json:"onlyZstack,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncEcsInstanceFromRemoteParam SyncEcsInstanceFromRemote request param
type SyncEcsInstanceFromRemoteParam struct {
	BaseParam
	Params SyncEcsInstanceFromRemoteDetailParam `json:"params"`
}
