// Copyright (c) ZStack.io, Inc.

package param

// SetVmSoundTypeDetailParam SetVmSoundType detail param
type SetVmSoundTypeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SoundType string `json:"soundType" validate:"required"`
}

// SetVmSoundTypeParam SetVmSoundType request param
type SetVmSoundTypeParam struct {
	BaseParam
	Params SetVmSoundTypeDetailParam `json:"params"`
}
