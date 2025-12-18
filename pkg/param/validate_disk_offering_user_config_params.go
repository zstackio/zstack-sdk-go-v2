// Copyright (c) ZStack.io, Inc.

package param

// ValidateDiskOfferingUserConfigDetailParam ValidateDiskOfferingUserConfig详细参数
type ValidateDiskOfferingUserConfigDetailParam struct {
	rest string `json:"config" validate:"required"` // 必填
}

// ValidateDiskOfferingUserConfigParam ValidateDiskOfferingUserConfig请求参数
type ValidateDiskOfferingUserConfigParam struct {
	BaseParam
	Params ValidateDiskOfferingUserConfigDetailParam `json:"params"` // 详细参数
}

