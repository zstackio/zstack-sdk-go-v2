// Copyright (c) ZStack.io, Inc.

package param

// DeleteGuestVmScriptDetailParam DeleteGuestVmScript detail param
type DeleteGuestVmScriptDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteGuestVmScriptParam DeleteGuestVmScript request param
type DeleteGuestVmScriptParam struct {
	BaseParam
	Params DeleteGuestVmScriptDetailParam `json:"params"`
}
