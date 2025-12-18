// Copyright (c) ZStack.io, Inc.

package param

// GetLoginProceduresDetailParam GetLoginProcedures detail param
type GetLoginProceduresDetailParam struct {
	Username string `json:"username" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
}

// GetLoginProceduresParam GetLoginProcedures request param
type GetLoginProceduresParam struct {
	BaseParam
	Params GetLoginProceduresDetailParam `json:"params"`
}
