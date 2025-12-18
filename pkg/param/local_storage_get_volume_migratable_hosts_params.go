// Copyright (c) ZStack.io, Inc.

package param

// LocalStorageGetVolumeMigratableHostsDetailParam LocalStorageGetVolumeMigratableHosts详细参数
type LocalStorageGetVolumeMigratableHostsDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
}

// LocalStorageGetVolumeMigratableHostsParam LocalStorageGetVolumeMigratableHosts请求参数
type LocalStorageGetVolumeMigratableHostsParam struct {
	BaseParam
	Params LocalStorageGetVolumeMigratableHostsDetailParam `json:"params"` // 详细参数
}

