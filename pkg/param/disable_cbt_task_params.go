// Copyright (c) ZStack.io, Inc.

package param

// DisableCbtTaskDetailParam DisableCbtTask detail param
type DisableCbtTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Force bool `json:"force,omitempty"`
}

// DisableCbtTaskParam DisableCbtTask request param
type DisableCbtTaskParam struct {
	BaseParam
	Params DisableCbtTaskDetailParam `json:"params"`
}
