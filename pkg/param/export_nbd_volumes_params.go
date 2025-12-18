// Copyright (c) ZStack.io, Inc.

package param

// ExportNbdVolumesDetailParam ExportNbdVolumes详细参数
type ExportNbdVolumesDetailParam struct {
	rest []string `json:"volumeUuids" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest bool `json:"force,omitempty"`
}

// ExportNbdVolumesParam ExportNbdVolumes请求参数
type ExportNbdVolumesParam struct {
	BaseParam
	Params ExportNbdVolumesDetailParam `json:"params"` // 详细参数
}

