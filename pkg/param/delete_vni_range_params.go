// Copyright (c) ZStack.io, Inc.

package param

// DeleteVniRangeDetailParam DeleteVniRange detail param
type DeleteVniRangeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVniRangeParam DeleteVniRange request param
type DeleteVniRangeParam struct {
	BaseParam
	Params DeleteVniRangeDetailParam `json:"params"`
}
