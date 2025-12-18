// Copyright (c) ZStack.io, Inc.

package param

// LogOutDetailParam LogOut detail param
type LogOutDetailParam struct {
	SessionUuid string `json:"sessionUuid,omitempty"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LogOutParam LogOut request param
type LogOutParam struct {
	BaseParam
	Params LogOutDetailParam `json:"params"`
}
