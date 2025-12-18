// Copyright (c) ZStack.io, Inc.

package param

// AttachNicToBondingDetailParam AttachNicToBonding detail param
type AttachNicToBondingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SlaveUuids []string `json:"slaveUuids" validate:"required"`
	Type string `json:"type,omitempty"`
}

// AttachNicToBondingParam AttachNicToBonding request param
type AttachNicToBondingParam struct {
	BaseParam
	Params AttachNicToBondingDetailParam `json:"params"`
}
