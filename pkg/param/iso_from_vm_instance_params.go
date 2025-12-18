// Copyright (c) ZStack.io, Inc.

package param

// DetachIsoFromVmInstanceDetailParam DetachIsoFromVmInstance详细参数
type DetachIsoFromVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"isoUuid,omitempty"`
}

// DetachIsoFromVmInstanceParam DetachIsoFromVmInstance请求参数
type DetachIsoFromVmInstanceParam struct {
	BaseParam
	Params DetachIsoFromVmInstanceDetailParam `json:"params"` // 详细参数
}

