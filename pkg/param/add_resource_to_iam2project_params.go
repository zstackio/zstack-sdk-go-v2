// Copyright (c) ZStack.io, Inc.

package param

// AddResourceToIAM2ProjectDetailParam AddResourceToIAM2Project详细参数
type AddResourceToIAM2ProjectDetailParam struct {
	rest string `json:"projectUuid" validate:"required"` // 必填
	rest []string `json:"resourceTemplates" validate:"required"` // 必填
}

// AddResourceToIAM2ProjectParam AddResourceToIAM2Project请求参数
type AddResourceToIAM2ProjectParam struct {
	BaseParam
	Params AddResourceToIAM2ProjectDetailParam `json:"params"` // 详细参数
}

