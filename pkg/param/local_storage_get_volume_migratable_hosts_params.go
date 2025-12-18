// Copyright (c) ZStack.io, Inc.

package param

// LocalStorageGetVolumeMigratableHostsDetailParam LocalStorageGetVolumeMigratableHosts detail param
type LocalStorageGetVolumeMigratableHostsDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
}

// LocalStorageGetVolumeMigratableHostsParam LocalStorageGetVolumeMigratableHosts request param
type LocalStorageGetVolumeMigratableHostsParam struct {
	BaseParam
	Params LocalStorageGetVolumeMigratableHostsDetailParam `json:"params"`
}
