// Copyright (c) ZStack.io, Inc.

package param

// DeleteEcsInstanceLocalDetailParam DeleteEcsInstanceLocal detail param
type DeleteEcsInstanceLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteEcsInstanceLocalParam DeleteEcsInstanceLocal request param
type DeleteEcsInstanceLocalParam struct {
	BaseParam
	Params DeleteEcsInstanceLocalDetailParam `json:"params"`
}
