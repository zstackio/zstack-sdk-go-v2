// Copyright (c) ZStack.io, Inc.

package param

// DeleteTagDetailParam DeleteTag detail param
type DeleteTagDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteTagParam DeleteTag request param
type DeleteTagParam struct {
	BaseParam
	Params DeleteTagDetailParam `json:"params"`
}
