// Copyright (c) ZStack.io, Inc.

package param

// DetachAliyunKeyDetailParam DetachAliyunKey detail param
type DetachAliyunKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DetachAliyunKeyParam DetachAliyunKey request param
type DetachAliyunKeyParam struct {
	BaseParam
	Params DetachAliyunKeyDetailParam `json:"params"`
}
