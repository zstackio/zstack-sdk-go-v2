// Copyright (c) ZStack.io, Inc.

package param

// DeletePublishAppDetailParam DeletePublishApp detail param
type DeletePublishAppDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePublishAppParam DeletePublishApp request param
type DeletePublishAppParam struct {
	BaseParam
	Params DeletePublishAppDetailParam `json:"params"`
}
