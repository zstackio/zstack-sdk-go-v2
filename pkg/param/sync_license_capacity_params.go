// Copyright (c) ZStack.io, Inc.

package param

// SyncLicenseCapacityDetailParam SyncLicenseCapacity detail param
type SyncLicenseCapacityDetailParam struct {
}

// SyncLicenseCapacityParam SyncLicenseCapacity request param
type SyncLicenseCapacityParam struct {
	BaseParam
	Params SyncLicenseCapacityDetailParam `json:"params"`
}
