// Copyright (c) ZStack.io, Inc.

package param

// GetUploadImageJobDetailsDetailParam GetUploadImageJobDetails detail param
type GetUploadImageJobDetailsDetailParam struct {
	ImageId string `json:"imageId" validate:"required"`
}

// GetUploadImageJobDetailsParam GetUploadImageJobDetails request param
type GetUploadImageJobDetailsParam struct {
	BaseParam
	Params GetUploadImageJobDetailsDetailParam `json:"params"`
}
