// Copyright (c) ZStack.io, Inc.

package param

// FlattenVolumeDetailParam FlattenVolume detail param
type FlattenVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DryRun bool `json:"dryRun,omitempty"`
}

// FlattenVolumeParam FlattenVolume request param
type FlattenVolumeParam struct {
	BaseParam
	Params FlattenVolumeDetailParam `json:"params"`
}
