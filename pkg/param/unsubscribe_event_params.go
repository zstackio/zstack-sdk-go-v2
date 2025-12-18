// Copyright (c) ZStack.io, Inc.

package param

// UnsubscribeEventDetailParam UnsubscribeEvent detail param
type UnsubscribeEventDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// UnsubscribeEventParam UnsubscribeEvent request param
type UnsubscribeEventParam struct {
	BaseParam
	Params UnsubscribeEventDetailParam `json:"params"`
}
