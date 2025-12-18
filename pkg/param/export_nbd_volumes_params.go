// Copyright (c) ZStack.io, Inc.

package param

// ExportNbdVolumesDetailParam ExportNbdVolumes detail param
type ExportNbdVolumesDetailParam struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Force bool `json:"force,omitempty"`
}

// ExportNbdVolumesParam ExportNbdVolumes request param
type ExportNbdVolumesParam struct {
	BaseParam
	Params ExportNbdVolumesDetailParam `json:"params"`
}
