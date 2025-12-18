// Copyright (c) ZStack.io, Inc.

package param

// UpdateVolumeDetailParam UpdateVolume detail param
type UpdateVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVolumeParam UpdateVolume request param
type UpdateVolumeParam struct {
	BaseParam
	Params UpdateVolumeDetailParam `json:"params"`
}
