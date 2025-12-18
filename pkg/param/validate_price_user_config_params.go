// Copyright (c) ZStack.io, Inc.

package param

// ValidatePriceUserConfigDetailParam ValidatePriceUserConfig详细参数
type ValidatePriceUserConfigDetailParam struct {
	rest string `json:"config" validate:"required"` // 必填
}

// ValidatePriceUserConfigParam ValidatePriceUserConfig请求参数
type ValidatePriceUserConfigParam struct {
	BaseParam
	Params ValidatePriceUserConfigDetailParam `json:"params"` // 详细参数
}

