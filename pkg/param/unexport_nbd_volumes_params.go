// Copyright (c) ZStack.io, Inc.

package param

// UnexportNbdVolumesDetailParam UnexportNbdVolumes detail param
type UnexportNbdVolumesDetailParam struct {
	VolumeUuids []string `json:"volumeUuids" validate:"required"`
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// UnexportNbdVolumesParam UnexportNbdVolumes request param
type UnexportNbdVolumesParam struct {
	BaseParam
	Params UnexportNbdVolumesDetailParam `json:"params"`
}
