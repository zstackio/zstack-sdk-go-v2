// Copyright (c) ZStack.io, Inc.

package param

// DeleteVolumeQosDetailParam DeleteVolumeQos detail param
type DeleteVolumeQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Mode string `json:"mode,omitempty"`
}

// DeleteVolumeQosParam DeleteVolumeQos request param
type DeleteVolumeQosParam struct {
	BaseParam
	Params DeleteVolumeQosDetailParam `json:"params"`
}
