// Copyright (c) ZStack.io, Inc.

package param

// UpdateGuestVmScriptDetailParam UpdateGuestVmScript detail param
type UpdateGuestVmScriptDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	EncodingType string `json:"encodingType,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
	RenderParams string `json:"renderParams,omitempty"`
	Platform string `json:"platform,omitempty"`
	ScriptType string `json:"scriptType,omitempty"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
}

// UpdateGuestVmScriptParam UpdateGuestVmScript request param
type UpdateGuestVmScriptParam struct {
	BaseParam
	Params UpdateGuestVmScriptDetailParam `json:"params"`
}
