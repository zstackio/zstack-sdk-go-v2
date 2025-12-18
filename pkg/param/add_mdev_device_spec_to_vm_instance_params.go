// Copyright (c) ZStack.io, Inc.

package param

// AddMdevDeviceSpecToVmInstanceDetailParam AddMdevDeviceSpecToVmInstance详细参数
type AddMdevDeviceSpecToVmInstanceDetailParam struct {
	rest string `json:"mdevSpecUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int `json:"mdevDeviceNumber,omitempty"`
}

// AddMdevDeviceSpecToVmInstanceParam AddMdevDeviceSpecToVmInstance请求参数
type AddMdevDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddMdevDeviceSpecToVmInstanceDetailParam `json:"params"` // 详细参数
}

