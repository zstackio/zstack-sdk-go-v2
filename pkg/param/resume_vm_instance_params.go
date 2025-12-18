// Copyright (c) ZStack.io, Inc.

package param

// ResumeVmInstanceDetailParam ResumeVmInstance详细参数
type ResumeVmInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ResumeVmInstanceParam ResumeVmInstance请求参数
type ResumeVmInstanceParam struct {
	BaseParam
	Params ResumeVmInstanceDetailParam `json:"params"` // 详细参数
}

