// Copyright (c) ZStack.io, Inc.

package param

// UpdateDiskOfferingDetailParam UpdateDiskOffering detail param
type UpdateDiskOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateDiskOfferingParam UpdateDiskOffering request param
type UpdateDiskOfferingParam struct {
	BaseParam
	Params UpdateDiskOfferingDetailParam `json:"params"`
}
