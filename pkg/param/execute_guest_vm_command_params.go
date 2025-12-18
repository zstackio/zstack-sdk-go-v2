// Copyright (c) ZStack.io, Inc.

package param

// ExecuteGuestVmCommandDetailParam ExecuteGuestVmCommand详细参数
type ExecuteGuestVmCommandDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"platform" validate:"required"` // 必填
	rest string `json:"command" validate:"required"` // 必填
	rest int `json:"commandTimeout,omitempty"`
}

// ExecuteGuestVmCommandParam ExecuteGuestVmCommand请求参数
type ExecuteGuestVmCommandParam struct {
	BaseParam
	Params ExecuteGuestVmCommandDetailParam `json:"params"` // 详细参数
}

