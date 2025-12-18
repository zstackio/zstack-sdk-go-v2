// Copyright (c) ZStack.io, Inc.

package param

// RemoveVmFromAffinityGroupDetailParam RemoveVmFromAffinityGroup detail param
type RemoveVmFromAffinityGroupDetailParam struct {
	AffinityGroupUuid string `json:"affinityGroupUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
}

// RemoveVmFromAffinityGroupParam RemoveVmFromAffinityGroup request param
type RemoveVmFromAffinityGroupParam struct {
	BaseParam
	Params RemoveVmFromAffinityGroupDetailParam `json:"params"`
}
