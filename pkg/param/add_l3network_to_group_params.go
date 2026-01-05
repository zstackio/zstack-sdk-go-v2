// Copyright (c) ZStack.io, Inc.

package param

// AddL3NetworkToGroupDetailParam AddL3NetworkToGroup detail param
type AddL3NetworkToGroupDetailParam struct {
	NfvInstGroupUuid string `json:"nfvInstGroupUuid" validate:"required"`
	NetworkServiceUuid string `json:"networkServiceUuid" validate:"required"`
	FrontEndL3NetworkUuid string `json:"frontEndL3NetworkUuid" validate:"required"`
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids" validate:"required"`
}

// AddL3NetworkToGroupParam AddL3NetworkToGroup request param
type AddL3NetworkToGroupParam struct {
	BaseParam
	Params AddL3NetworkToGroupDetailParam `json:"params"`
}
