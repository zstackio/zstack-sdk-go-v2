// Copyright (c) ZStack.io, Inc.

package param

// ExpungeDataVolumeDetailParam ExpungeDataVolume detail param
type ExpungeDataVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeDataVolumeParam ExpungeDataVolume request param
type ExpungeDataVolumeParam struct {
	BaseParam
	Params ExpungeDataVolumeDetailParam `json:"params"`
}
