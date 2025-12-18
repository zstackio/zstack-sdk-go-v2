// Copyright (c) ZStack.io, Inc.

package param

// UnexportNbdVolumesDetailParam UnexportNbdVolumes详细参数
type UnexportNbdVolumesDetailParam struct {
	rest []string `json:"volumeUuids" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// UnexportNbdVolumesParam UnexportNbdVolumes请求参数
type UnexportNbdVolumesParam struct {
	BaseParam
	Params UnexportNbdVolumesDetailParam `json:"params"` // 详细参数
}

