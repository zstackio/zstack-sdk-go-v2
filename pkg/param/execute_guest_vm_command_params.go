// Copyright (c) ZStack.io, Inc.

package param

// ExecuteGuestVmCommandDetailParam ExecuteGuestVmCommand detail param
type ExecuteGuestVmCommandDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	Platform string `json:"platform" validate:"required"`
	Command string `json:"command" validate:"required"`
	CommandTimeout int `json:"commandTimeout,omitempty"`
}

// ExecuteGuestVmCommandParam ExecuteGuestVmCommand request param
type ExecuteGuestVmCommandParam struct {
	BaseParam
	Params ExecuteGuestVmCommandDetailParam `json:"params"`
}
