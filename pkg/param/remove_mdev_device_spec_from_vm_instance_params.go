// Copyright (c) ZStack.io, Inc.

package param

// RemoveMdevDeviceSpecFromVmInstanceDetailParam RemoveMdevDeviceSpecFromVmInstance详细参数
type RemoveMdevDeviceSpecFromVmInstanceDetailParam struct {
	rest string `json:"mdevSpecUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// RemoveMdevDeviceSpecFromVmInstanceParam RemoveMdevDeviceSpecFromVmInstance请求参数
type RemoveMdevDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemoveMdevDeviceSpecFromVmInstanceDetailParam `json:"params"` // 详细参数
}

