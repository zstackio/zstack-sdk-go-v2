// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateMdevDeviceParamDetail UpdateMdevDevice detail param
type UpdateMdevDeviceParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateMdevDeviceParam UpdateMdevDevice request param
type UpdateMdevDeviceParam struct {
	BaseParam
	UpdateMdevDevice UpdateMdevDeviceParamDetail `json:"updateMdevDevice"`
}
// DeleteMdevDeviceParamDetail DeleteMdevDevice detail param
type DeleteMdevDeviceParamDetail struct {
	MdevDeviceUuid string `json:"mdevDeviceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMdevDeviceParam DeleteMdevDevice request param
type DeleteMdevDeviceParam struct {
	BaseParam
	DeleteMdevDevice DeleteMdevDeviceParamDetail `json:"deleteMdevDevice"`
}
