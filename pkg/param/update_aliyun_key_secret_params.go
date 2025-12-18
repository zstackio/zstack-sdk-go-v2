// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunKeySecretDetailParam UpdateAliyunKeySecret detail param
type UpdateAliyunKeySecretDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunKeySecretParam UpdateAliyunKeySecret request param
type UpdateAliyunKeySecretParam struct {
	BaseParam
	Params UpdateAliyunKeySecretDetailParam `json:"params"`
}
