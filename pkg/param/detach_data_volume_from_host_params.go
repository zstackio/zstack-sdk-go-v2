// Copyright (c) ZStack.io, Inc.

package param

// DetachDataVolumeFromHostDetailParam DetachDataVolumeFromHost detail param
type DetachDataVolumeFromHostDetailParam struct {
	VolumeUuid string `json:"volumeUuid" validate:"required"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// DetachDataVolumeFromHostParam DetachDataVolumeFromHost request param
type DetachDataVolumeFromHostParam struct {
	BaseParam
	Params DetachDataVolumeFromHostDetailParam `json:"params"`
}
