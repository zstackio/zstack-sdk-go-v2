// Copyright (c) ZStack.io, Inc.

package param

// CreateGuestVmScriptDetailParam CreateGuestVmScript detail param
type CreateGuestVmScriptDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	EncodingType string `json:"encodingType" validate:"required"`
	ScriptContent string `json:"scriptContent" validate:"required"`
	RenderParams string `json:"renderParams,omitempty"`
	Platform string `json:"platform" validate:"required"`
	ScriptType string `json:"scriptType" validate:"required"`
	ScriptTimeout int `json:"scriptTimeout,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateGuestVmScriptParam CreateGuestVmScript request param
type CreateGuestVmScriptParam struct {
	BaseParam
	Params CreateGuestVmScriptDetailParam `json:"params"`
}
