// Copyright (c) ZStack.io, Inc.

package param

// DetachNfvInstFromGroupDetailParam DetachNfvInstFromGroup detail param
type DetachNfvInstFromGroupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	NfvInstUuid string `json:"nfvInstUuid" validate:"required"`
}

// DetachNfvInstFromGroupParam DetachNfvInstFromGroup request param
type DetachNfvInstFromGroupParam struct {
	BaseParam
	Params DetachNfvInstFromGroupDetailParam `json:"params"`
}
