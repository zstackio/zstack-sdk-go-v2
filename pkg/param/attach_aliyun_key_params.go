// Copyright (c) ZStack.io, Inc.

package param

// AttachAliyunKeyDetailParam AttachAliyunKey detail param
type AttachAliyunKeyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// AttachAliyunKeyParam AttachAliyunKey request param
type AttachAliyunKeyParam struct {
	BaseParam
	Params AttachAliyunKeyDetailParam `json:"params"`
}
