// Copyright (c) ZStack.io, Inc.

package param

// ExecuteGuestVmScriptDetailParam ExecuteGuestVmScript detail param
type ExecuteGuestVmScriptDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmInstanceUuids []string `json:"vmInstanceUuids" validate:"required"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
	LogPath string `json:"logPath,omitempty"`
	RecordUuid string `json:"recordUuid,omitempty"`
	RuntimeParams string `json:"runtimeParams,omitempty"`
}

// ExecuteGuestVmScriptParam ExecuteGuestVmScript request param
type ExecuteGuestVmScriptParam struct {
	BaseParam
	Params ExecuteGuestVmScriptDetailParam `json:"params"`
}
