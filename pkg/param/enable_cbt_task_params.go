// Copyright (c) ZStack.io, Inc.

package param

// EnableCbtTaskDetailParam EnableCbtTask detail param
type EnableCbtTaskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BitmapName string `json:"bitmapName,omitempty"`
}

// EnableCbtTaskParam EnableCbtTask request param
type EnableCbtTaskParam struct {
	BaseParam
	Params EnableCbtTaskDetailParam `json:"params"`
}
