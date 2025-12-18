// Copyright (c) ZStack.io, Inc.

package param

// LocateLocalRaidPhysicalDriveDetailParam LocateLocalRaidPhysicalDrive详细参数
type LocateLocalRaidPhysicalDriveDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"locate,omitempty"`
}

// LocateLocalRaidPhysicalDriveParam LocateLocalRaidPhysicalDrive请求参数
type LocateLocalRaidPhysicalDriveParam struct {
	BaseParam
	Params LocateLocalRaidPhysicalDriveDetailParam `json:"params"` // 详细参数
}

