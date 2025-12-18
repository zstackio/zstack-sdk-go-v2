// Copyright (c) ZStack.io, Inc.

package param

// GetCreateEcsImageProgressDetailParam GetCreateEcsImageProgress detail param
type GetCreateEcsImageProgressDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
}

// GetCreateEcsImageProgressParam GetCreateEcsImageProgress request param
type GetCreateEcsImageProgressParam struct {
	BaseParam
	Params GetCreateEcsImageProgressDetailParam `json:"params"`
}
