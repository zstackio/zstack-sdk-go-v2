// Copyright (c) ZStack.io, Inc.

package param

// RestartResourceStackDetailParam RestartResourceStack detail param
type RestartResourceStackDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// RestartResourceStackParam RestartResourceStack request param
type RestartResourceStackParam struct {
	BaseParam
	Params RestartResourceStackDetailParam `json:"params"`
}
