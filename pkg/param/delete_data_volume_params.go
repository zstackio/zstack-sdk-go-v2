// Copyright (c) ZStack.io, Inc.

package param

// DeleteDataVolumeDetailParam DeleteDataVolume detail param
type DeleteDataVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDataVolumeParam DeleteDataVolume request param
type DeleteDataVolumeParam struct {
	BaseParam
	Params DeleteDataVolumeDetailParam `json:"params"`
}
