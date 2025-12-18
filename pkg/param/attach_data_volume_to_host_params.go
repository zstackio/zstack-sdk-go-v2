// Copyright (c) ZStack.io, Inc.

package param

// AttachDataVolumeToHostDetailParam AttachDataVolumeToHost detail param
type AttachDataVolumeToHostDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	MountPath string `json:"mountPath" validate:"required"`
}

// AttachDataVolumeToHostParam AttachDataVolumeToHost request param
type AttachDataVolumeToHostParam struct {
	BaseParam
	Params AttachDataVolumeToHostDetailParam `json:"params"`
}
