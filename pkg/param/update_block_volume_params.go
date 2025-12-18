// Copyright (c) ZStack.io, Inc.

package param

// UpdateBlockVolumeDetailParam UpdateBlockVolume detail param
type UpdateBlockVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBlockVolumeParam UpdateBlockVolume request param
type UpdateBlockVolumeParam struct {
	BaseParam
	Params UpdateBlockVolumeDetailParam `json:"params"`
}
