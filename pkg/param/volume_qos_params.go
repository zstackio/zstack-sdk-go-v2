// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeQosDetailParam GetVolumeQos详细参数
type GetVolumeQosDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"forceSync,omitempty"`
}

// GetVolumeQosParam GetVolumeQos请求参数
type GetVolumeQosParam struct {
	BaseParam
	Params GetVolumeQosDetailParam `json:"params"` // 详细参数
}

