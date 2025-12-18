// Copyright (c) ZStack.io, Inc.

package param

// ResumeVmInstanceDetailParam ResumeVmInstance detail param
type ResumeVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ResumeVmInstanceParam ResumeVmInstance request param
type ResumeVmInstanceParam struct {
	BaseParam
	Params ResumeVmInstanceDetailParam `json:"params"`
}
