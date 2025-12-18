// Copyright (c) ZStack.io, Inc.

package param

// UpdateScsiLunDetailParam UpdateScsiLun detail param
type UpdateScsiLunDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateScsiLunParam UpdateScsiLun request param
type UpdateScsiLunParam struct {
	BaseParam
	Params UpdateScsiLunDetailParam `json:"params"`
}
