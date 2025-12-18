// Copyright (c) ZStack.io, Inc.

package param

// DeleteConnectionAccessPointLocalDetailParam DeleteConnectionAccessPointLocal detail param
type DeleteConnectionAccessPointLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteConnectionAccessPointLocalParam DeleteConnectionAccessPointLocal request param
type DeleteConnectionAccessPointLocalParam struct {
	BaseParam
	Params DeleteConnectionAccessPointLocalDetailParam `json:"params"`
}
