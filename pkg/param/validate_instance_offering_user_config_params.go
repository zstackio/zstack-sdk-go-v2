// Copyright (c) ZStack.io, Inc.

package param

// ValidateInstanceOfferingUserConfigDetailParam ValidateInstanceOfferingUserConfig详细参数
type ValidateInstanceOfferingUserConfigDetailParam struct {
	rest string `json:"config" validate:"required"` // 必填
}

// ValidateInstanceOfferingUserConfigParam ValidateInstanceOfferingUserConfig请求参数
type ValidateInstanceOfferingUserConfigParam struct {
	BaseParam
	Params ValidateInstanceOfferingUserConfigDetailParam `json:"params"` // 详细参数
}

