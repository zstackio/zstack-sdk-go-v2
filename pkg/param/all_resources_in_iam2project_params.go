// Copyright (c) ZStack.io, Inc.

package param

// StopAllResourcesInIAM2ProjectDetailParam StopAllResourcesInIAM2Project详细参数
type StopAllResourcesInIAM2ProjectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// StopAllResourcesInIAM2ProjectParam StopAllResourcesInIAM2Project请求参数
type StopAllResourcesInIAM2ProjectParam struct {
	BaseParam
	Params StopAllResourcesInIAM2ProjectDetailParam `json:"params"` // 详细参数
}

