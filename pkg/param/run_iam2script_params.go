// Copyright (c) ZStack.io, Inc.

package param

// RunIAM2ScriptDetailParam RunIAM2Script detail param
type RunIAM2ScriptDetailParam struct {
	ScriptContent string `json:"scriptContent" validate:"required"`
	ScriptExecutor string `json:"scriptExecutor,omitempty"`
	ScriptParams []string `json:"scriptParams,omitempty"`
}

// RunIAM2ScriptParam RunIAM2Script request param
type RunIAM2ScriptParam struct {
	BaseParam
	Params RunIAM2ScriptDetailParam `json:"params"`
}
