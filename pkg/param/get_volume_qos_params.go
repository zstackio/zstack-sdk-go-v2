// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeQosDetailParam GetVolumeQos detail param
type GetVolumeQosDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ForceSync bool `json:"forceSync,omitempty"`
}

// GetVolumeQosParam GetVolumeQos request param
type GetVolumeQosParam struct {
	BaseParam
	Params GetVolumeQosDetailParam `json:"params"`
}
