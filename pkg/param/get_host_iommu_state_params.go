// Copyright (c) ZStack.io, Inc.

package param

// GetHostIommuStateDetailParam GetHostIommuState detail param
type GetHostIommuStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetHostIommuStateParam GetHostIommuState request param
type GetHostIommuStateParam struct {
	BaseParam
	Params GetHostIommuStateDetailParam `json:"params"`
}
