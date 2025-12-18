// Copyright (c) ZStack.io, Inc.

package param

// ExecuteGuestVmScriptDetailParam ExecuteGuestVmScript详细参数
type ExecuteGuestVmScriptDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"vmInstanceUuids" validate:"required"` // 必填
	rest int `json:"scriptTimeout,omitempty"`
	rest string `json:"logPath,omitempty"`
	rest string `json:"recordUuid,omitempty"`
}

// ExecuteGuestVmScriptParam ExecuteGuestVmScript请求参数
type ExecuteGuestVmScriptParam struct {
	BaseParam
	Params ExecuteGuestVmScriptDetailParam `json:"params"` // 详细参数
}

