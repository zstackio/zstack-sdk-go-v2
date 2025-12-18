// Copyright (c) ZStack.io, Inc.

package param

// LoginByCasDetailParam LoginByCas detail param
type LoginByCasDetailParam struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type string `json:"type" validate:"required"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginByCasParam LoginByCas request param
type LoginByCasParam struct {
	BaseParam
	Params LoginByCasDetailParam `json:"params"`
}
