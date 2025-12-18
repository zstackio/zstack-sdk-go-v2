// Copyright (c) ZStack.io, Inc.

package param

// DetachScsiLunFromVmInstanceDetailParam DetachScsiLunFromVmInstance详细参数
type DetachScsiLunFromVmInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// DetachScsiLunFromVmInstanceParam DetachScsiLunFromVmInstance请求参数
type DetachScsiLunFromVmInstanceParam struct {
	BaseParam
	Params DetachScsiLunFromVmInstanceDetailParam `json:"params"` // 详细参数
}

