// Copyright (c) ZStack.io, Inc.

package param

// AttachScsiLunToVmInstanceDetailParam AttachScsiLunToVmInstance详细参数
type AttachScsiLunToVmInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest bool `json:"disableMultiPathAttach,omitempty"`
}

// AttachScsiLunToVmInstanceParam AttachScsiLunToVmInstance请求参数
type AttachScsiLunToVmInstanceParam struct {
	BaseParam
	Params AttachScsiLunToVmInstanceDetailParam `json:"params"` // 详细参数
}

