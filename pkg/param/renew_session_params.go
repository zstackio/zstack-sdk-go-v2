// Copyright (c) ZStack.io, Inc.

package param

// RenewSessionDetailParam RenewSession detail param
type RenewSessionDetailParam struct {
	SessionUuid string `json:"sessionUuid" validate:"required"`
	Duration int64 `json:"duration,omitempty"`
}

// RenewSessionParam RenewSession request param
type RenewSessionParam struct {
	BaseParam
	Params RenewSessionDetailParam `json:"params"`
}
