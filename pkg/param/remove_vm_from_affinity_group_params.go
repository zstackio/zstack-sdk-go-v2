// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmFromAffinityGroupDetailParam RemoveVmFromAffinityGroup详细参数
type RemoveVmFromAffinityGroupDetailParam struct {
	rest string `json:"affinityGroupUuid" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
}

// RemoveVmFromAffinityGroupParam RemoveVmFromAffinityGroup请求参数
type RemoveVmFromAffinityGroupParam struct {
	BaseParam
	Params RemoveVmFromAffinityGroupDetailParam `json:"params"` // 详细参数
}

