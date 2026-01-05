// Copyright (c) ZStack.io, Inc.

package param

// AttachNfvInstToGroupDetailParam AttachNfvInstToGroup detail param
type AttachNfvInstToGroupDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	NfvInstUuid string `json:"nfvInstUuid" validate:"required"`
}

// AttachNfvInstToGroupParam AttachNfvInstToGroup request param
type AttachNfvInstToGroupParam struct {
	BaseParam
	Params AttachNfvInstToGroupDetailParam `json:"params"`
}
