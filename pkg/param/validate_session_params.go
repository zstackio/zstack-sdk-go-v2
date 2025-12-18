// Copyright (c) ZStack.io, Inc.

package param

// ValidateSessionDetailParam ValidateSession detail param
type ValidateSessionDetailParam struct {
	SessionUuid string `json:"sessionUuid" validate:"required"`
}

// ValidateSessionParam ValidateSession request param
type ValidateSessionParam struct {
	BaseParam
	Params ValidateSessionDetailParam `json:"params"`
}
