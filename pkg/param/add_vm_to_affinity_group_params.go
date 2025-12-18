// Copyright (c) ZStack.io, Inc.

package param

// AddVmToAffinityGroupDetailParam AddVmToAffinityGroup detail param
type AddVmToAffinityGroupDetailParam struct {
	AffinityGroupUuid string `json:"affinityGroupUuid" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
}

// AddVmToAffinityGroupParam AddVmToAffinityGroup request param
type AddVmToAffinityGroupParam struct {
	BaseParam
	Params AddVmToAffinityGroupDetailParam `json:"params"`
}
