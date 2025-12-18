// Copyright (c) ZStack.io, Inc.

package param

// GetLocalRaidPhysicalDriveSmartDetailParam GetLocalRaidPhysicalDriveSmart详细参数
type GetLocalRaidPhysicalDriveSmartDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetLocalRaidPhysicalDriveSmartParam GetLocalRaidPhysicalDriveSmart请求参数
type GetLocalRaidPhysicalDriveSmartParam struct {
	BaseParam
	Params GetLocalRaidPhysicalDriveSmartDetailParam `json:"params"` // 详细参数
}

