// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostIommuStateDetailParam UpdateHostIommuState detail param
type UpdateHostIommuStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	State string `json:"state" validate:"required"`
}

// UpdateHostIommuStateParam UpdateHostIommuState request param
type UpdateHostIommuStateParam struct {
	BaseParam
	Params UpdateHostIommuStateDetailParam `json:"params"`
}
