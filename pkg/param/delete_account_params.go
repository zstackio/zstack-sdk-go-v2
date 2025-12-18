// Copyright (c) ZStack.io, Inc.

package param

// DeleteAccountDetailParam DeleteAccount detail param
type DeleteAccountDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAccountParam DeleteAccount request param
type DeleteAccountParam struct {
	BaseParam
	Params DeleteAccountDetailParam `json:"params"`
}
