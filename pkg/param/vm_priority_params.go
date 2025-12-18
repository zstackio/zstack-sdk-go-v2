// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmPriorityDetailParam UpdateVmPriority详细参数
type UpdateVmPriorityDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"priority" validate:"required"` // 必填
}

// UpdateVmPriorityParam UpdateVmPriority请求参数
type UpdateVmPriorityParam struct {
	BaseParam
	Params UpdateVmPriorityDetailParam `json:"params"` // 详细参数
}

