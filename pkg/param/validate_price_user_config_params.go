// Copyright (c) ZStack.io, Inc.

package param

// ValidatePriceUserConfigDetailParam ValidatePriceUserConfig detail param
type ValidatePriceUserConfigDetailParam struct {
	Config string `json:"config" validate:"required"`
}

// ValidatePriceUserConfigParam ValidatePriceUserConfig request param
type ValidatePriceUserConfigParam struct {
	BaseParam
	Params ValidatePriceUserConfigDetailParam `json:"params"`
}
