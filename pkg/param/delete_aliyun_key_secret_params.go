// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunKeySecretDetailParam DeleteAliyunKeySecret detail param
type DeleteAliyunKeySecretDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunKeySecretParam DeleteAliyunKeySecret request param
type DeleteAliyunKeySecretParam struct {
	BaseParam
	Params DeleteAliyunKeySecretDetailParam `json:"params"`
}
