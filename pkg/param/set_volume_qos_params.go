// Copyright (c) ZStack.io, Inc.

package param

// SetVolumeQosDetailParam SetVolumeQos detail param
type SetVolumeQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Mode string `json:"mode,omitempty"`
	VolumeBandwidth int64 `json:"volumeBandwidth,omitempty"`
	ReadBandwidth int64 `json:"readBandwidth,omitempty"`
	WriteBandwidth int64 `json:"writeBandwidth,omitempty"`
	TotalBandwidth int64 `json:"totalBandwidth,omitempty"`
	ReadIOPS int64 `json:"readIOPS,omitempty"`
	WriteIOPS int64 `json:"writeIOPS,omitempty"`
	TotalIOPS int64 `json:"totalIOPS,omitempty"`
}

// SetVolumeQosParam SetVolumeQos request param
type SetVolumeQosParam struct {
	BaseParam
	Params SetVolumeQosDetailParam `json:"params"`
}
