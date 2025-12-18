// Copyright (c) ZStack.io, Inc.

package param

// GetVmAttachableL3NetworkDetailParam GetVmAttachableL3Network详细参数
type GetVmAttachableL3NetworkDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetVmAttachableL3NetworkParam GetVmAttachableL3Network请求参数
type GetVmAttachableL3NetworkParam struct {
	BaseParam
	Params GetVmAttachableL3NetworkDetailParam `json:"params"` // 详细参数
}

