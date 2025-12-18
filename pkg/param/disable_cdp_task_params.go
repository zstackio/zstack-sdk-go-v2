// Copyright (c) ZStack.io, Inc.

package param

// DisableCdpTaskDetailParam DisableCdpTask detail param
type DisableCdpTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
}

// DisableCdpTaskParam DisableCdpTask request param
type DisableCdpTaskParam struct {
	BaseParam
	Params DisableCdpTaskDetailParam `json:"params"`
}
