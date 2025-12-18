// Copyright (c) ZStack.io, Inc.

package param

// AddVmToAffinityGroupDetailParam AddVmToAffinityGroup详细参数
type AddVmToAffinityGroupDetailParam struct {
	rest string `json:"affinityGroupUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
}

// AddVmToAffinityGroupParam AddVmToAffinityGroup请求参数
type AddVmToAffinityGroupParam struct {
	BaseParam
	Params AddVmToAffinityGroupDetailParam `json:"params"` // 详细参数
}

