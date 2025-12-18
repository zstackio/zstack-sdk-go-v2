// Copyright (c) ZStack.io, Inc.

package param

// DeleteDiskOfferingDetailParam DeleteDiskOffering detail param
type DeleteDiskOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteDiskOfferingParam DeleteDiskOffering request param
type DeleteDiskOfferingParam struct {
	BaseParam
	Params DeleteDiskOfferingDetailParam `json:"params"`
}
