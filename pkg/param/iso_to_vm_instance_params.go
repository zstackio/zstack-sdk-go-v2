// Copyright (c) ZStack.io, Inc.

package param

// AttachIsoToVmInstanceDetailParam AttachIsoToVmInstance详细参数
type AttachIsoToVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"isoUuid" validate:"required"` // 必填
}

// AttachIsoToVmInstanceParam AttachIsoToVmInstance请求参数
type AttachIsoToVmInstanceParam struct {
	BaseParam
	Params AttachIsoToVmInstanceDetailParam `json:"params"` // 详细参数
}

