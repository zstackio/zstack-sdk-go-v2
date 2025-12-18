// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunDiskDetailParam UpdateAliyunDisk详细参数
type UpdateAliyunDiskDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"deleteWithInstance,omitempty"`
	rest bool `json:"deleteAutoSnapshot,omitempty"`
	rest bool `json:"enableAutoSnapshot,omitempty"`
}

// UpdateAliyunDiskParam UpdateAliyunDisk请求参数
type UpdateAliyunDiskParam struct {
	BaseParam
	Params UpdateAliyunDiskDetailParam `json:"params"` // 详细参数
}

