// Copyright (c) ZStack.io, Inc.

package param

// UpdateXskyBlockVolumeDetailParam UpdateXskyBlockVolume detail param
type UpdateXskyBlockVolumeDetailParam struct {
	BurstTotalBw int64 `json:"burstTotalBw,omitempty"`
	BurstTotalIops int64 `json:"burstTotalIops,omitempty"`
	MaxTotalBw int64 `json:"maxTotalBw,omitempty"`
	MaxTotalIops int64 `json:"maxTotalIops,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateXskyBlockVolumeParam UpdateXskyBlockVolume request param
type UpdateXskyBlockVolumeParam struct {
	BaseParam
	Params UpdateXskyBlockVolumeDetailParam `json:"params"`
}
