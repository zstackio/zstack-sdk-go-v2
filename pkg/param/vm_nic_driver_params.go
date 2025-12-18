// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmNicDriverDetailParam UpdateVmNicDriver详细参数
type UpdateVmNicDriverDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"driverType" validate:"required"` // 必填
}

// UpdateVmNicDriverParam UpdateVmNicDriver请求参数
type UpdateVmNicDriverParam struct {
	BaseParam
	Params UpdateVmNicDriverDetailParam `json:"params"` // 详细参数
}

