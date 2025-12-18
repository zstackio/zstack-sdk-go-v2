// Copyright (c) ZStack.io, Inc.

package param

// GetVipQosDetailParam GetVipQos详细参数
type GetVipQosDetailParam struct {
	rest string `json:"uuid,omitempty"`
}

// GetVipQosParam GetVipQos请求参数
type GetVipQosParam struct {
	BaseParam
	Params GetVipQosDetailParam `json:"params"` // 详细参数
}

