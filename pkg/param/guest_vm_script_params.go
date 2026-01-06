// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteGuestVmScriptParamDetail DeleteGuestVmScript detail param
type DeleteGuestVmScriptParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteGuestVmScriptParam DeleteGuestVmScript request param
type DeleteGuestVmScriptParam struct {
	BaseParam
	Params DeleteGuestVmScriptParamDetail `json:"params"`
}
// CreateGuestVmScriptParamDetail CreateGuestVmScript detail param
type CreateGuestVmScriptParamDetail struct {
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
	Params CreateGuestVmScriptParamDetail `json:"params"`
}
// UpdateGuestVmScriptParamDetail UpdateGuestVmScript detail param
type UpdateGuestVmScriptParamDetail struct {
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
	Params UpdateGuestVmScriptParamDetail `json:"params"`
}
