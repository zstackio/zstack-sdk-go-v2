// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourcePriceDetailParam UpdateResourcePrice detail param
type UpdateResourcePriceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	EndDateInLong int64 `json:"endDateInLong,omitempty"`
	SetEndDateInLongBaseOnCurrentTime bool `json:"setEndDateInLongBaseOnCurrentTime,omitempty"`
}

// UpdateResourcePriceParam UpdateResourcePrice request param
type UpdateResourcePriceParam struct {
	BaseParam
	Params UpdateResourcePriceDetailParam `json:"params"`
}
