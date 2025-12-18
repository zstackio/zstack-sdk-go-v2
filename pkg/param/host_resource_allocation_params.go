// Copyright (c) ZStack.io, Inc.

package param

// GetHostResourceAllocationDetailParam GetHostResourceAllocation详细参数
type GetHostResourceAllocationDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"strategy" validate:"required"` // 必填
	rest string `json:"scene" validate:"required"` // 必填
	rest int `json:"vcpu" validate:"required"` // 必填
	rest int64 `json:"memSize,omitempty"`
}

// GetHostResourceAllocationParam GetHostResourceAllocation请求参数
type GetHostResourceAllocationParam struct {
	BaseParam
	Params GetHostResourceAllocationDetailParam `json:"params"` // 详细参数
}

