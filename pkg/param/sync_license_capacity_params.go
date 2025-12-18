// Copyright (c) ZStack.io, Inc.

package param

// SyncLicenseCapacityDetailParam SyncLicenseCapacity详细参数
type SyncLicenseCapacityDetailParam struct {
}

// SyncLicenseCapacityParam SyncLicenseCapacity请求参数
type SyncLicenseCapacityParam struct {
	BaseParam
	Params SyncLicenseCapacityDetailParam `json:"params"` // 详细参数
}

