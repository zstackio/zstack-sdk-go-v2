// Copyright (c) ZStack.io, Inc.

package param

// SetIAM2ProjectRetirePolicyDetailParam SetIAM2ProjectRetirePolicy detail param
type SetIAM2ProjectRetirePolicyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Policy string `json:"policy" validate:"required"`
}

// SetIAM2ProjectRetirePolicyParam SetIAM2ProjectRetirePolicy request param
type SetIAM2ProjectRetirePolicyParam struct {
	BaseParam
	Params SetIAM2ProjectRetirePolicyDetailParam `json:"params"`
}
