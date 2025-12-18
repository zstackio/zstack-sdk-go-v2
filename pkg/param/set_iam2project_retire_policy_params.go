// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectRetirePolicyDetailParam SetIAM2ProjectRetirePolicy详细参数
type SetIAM2ProjectRetirePolicyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"policy" validate:"required"` // 必填
}

// SetIAM2ProjectRetirePolicyParam SetIAM2ProjectRetirePolicy请求参数
type SetIAM2ProjectRetirePolicyParam struct {
	BaseParam
	Params SetIAM2ProjectRetirePolicyDetailParam `json:"params"` // 详细参数
}

