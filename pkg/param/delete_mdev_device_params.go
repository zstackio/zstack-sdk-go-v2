// Copyright (c) ZStack.io, Inc.

package param

// DeleteMdevDeviceDetailParam DeleteMdevDevice detail param
type DeleteMdevDeviceDetailParam struct {
	MdevDeviceUuid string `json:"mdevDeviceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMdevDeviceParam DeleteMdevDevice request param
type DeleteMdevDeviceParam struct {
	BaseParam
	Params DeleteMdevDeviceDetailParam `json:"params"`
}
