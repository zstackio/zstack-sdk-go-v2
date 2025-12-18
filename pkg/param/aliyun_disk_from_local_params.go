// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunDiskFromLocalDetailParam DeleteAliyunDiskFromLocal详细参数
type DeleteAliyunDiskFromLocalDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAliyunDiskFromLocalParam DeleteAliyunDiskFromLocal请求参数
type DeleteAliyunDiskFromLocalParam struct {
	BaseParam
	Params DeleteAliyunDiskFromLocalDetailParam `json:"params"` // 详细参数
}

