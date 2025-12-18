// Copyright (c) ZStack.io, Inc.

package param

// DeleteCbtTaskDetailParam DeleteCbtTask detail param
type DeleteCbtTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteCbtTaskParam DeleteCbtTask request param
type DeleteCbtTaskParam struct {
	BaseParam
	Params DeleteCbtTaskDetailParam `json:"params"`
}
