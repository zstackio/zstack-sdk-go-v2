// Copyright (c) ZStack.io, Inc.

package param

// AttachHybridEipToEcsDetailParam AttachHybridEipToEcs详细参数
type AttachHybridEipToEcsDetailParam struct {
	rest string `json:"eipUuid" validate:"required"` // 必填
	rest string `json:"ecsUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
}

// AttachHybridEipToEcsParam AttachHybridEipToEcs请求参数
type AttachHybridEipToEcsParam struct {
	BaseParam
	Params AttachHybridEipToEcsDetailParam `json:"params"` // 详细参数
}

