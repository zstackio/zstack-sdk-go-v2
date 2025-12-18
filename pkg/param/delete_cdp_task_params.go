// Copyright (c) ZStack.io, Inc.

package param

// DeleteCdpTaskDetailParam DeleteCdpTask detail param
type DeleteCdpTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteCdpTaskParam DeleteCdpTask request param
type DeleteCdpTaskParam struct {
	BaseParam
	Params DeleteCdpTaskDetailParam `json:"params"`
}
