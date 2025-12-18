// Copyright (c) ZStack.io, Inc.

package param

// ReimageVmInstanceDetailParam ReimageVmInstance详细参数
type ReimageVmInstanceDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// ReimageVmInstanceParam ReimageVmInstance请求参数
type ReimageVmInstanceParam struct {
	BaseParam
	Params ReimageVmInstanceDetailParam `json:"params"` // 详细参数
}

