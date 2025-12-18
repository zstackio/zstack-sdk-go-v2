// Copyright (c) ZStack.io, Inc.

package param

// DetachHybridEipFromEcsDetailParam DetachHybridEipFromEcs详细参数
type DetachHybridEipFromEcsDetailParam struct {
	rest string `json:"eipUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
}

// DetachHybridEipFromEcsParam DetachHybridEipFromEcs请求参数
type DetachHybridEipFromEcsParam struct {
	BaseParam
	Params DetachHybridEipFromEcsDetailParam `json:"params"` // 详细参数
}

