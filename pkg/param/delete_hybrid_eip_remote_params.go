// Copyright (c) ZStack.io, Inc.

package param

// DeleteHybridEipRemoteDetailParam DeleteHybridEipRemote detail param
type DeleteHybridEipRemoteDetailParam struct {
	Type string `json:"type" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHybridEipRemoteParam DeleteHybridEipRemote request param
type DeleteHybridEipRemoteParam struct {
	BaseParam
	Params DeleteHybridEipRemoteDetailParam `json:"params"`
}
