// Copyright (c) ZStack.io, Inc.

package param

// CheckIAM2VirtualIDConfigFileDetailParam CheckIAM2VirtualIDConfigFile详细参数
type CheckIAM2VirtualIDConfigFileDetailParam struct {
	rest string `json:"virtualIDInfos" validate:"required"` // 必填
}

// CheckIAM2VirtualIDConfigFileParam CheckIAM2VirtualIDConfigFile请求参数
type CheckIAM2VirtualIDConfigFileParam struct {
	BaseParam
	Params CheckIAM2VirtualIDConfigFileDetailParam `json:"params"` // 详细参数
}

