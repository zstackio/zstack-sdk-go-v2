// Copyright (c) ZStack.io, Inc.

package param

// RecoverDataVolumeDetailParam RecoverDataVolume detail param
type RecoverDataVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RecoverDataVolumeParam RecoverDataVolume request param
type RecoverDataVolumeParam struct {
	BaseParam
	Params RecoverDataVolumeDetailParam `json:"params"`
}
