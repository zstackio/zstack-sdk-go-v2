// Copyright (c) ZStack.io, Inc.

package param

// TriggerGCJobDetailParam TriggerGCJob详细参数
type TriggerGCJobDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// TriggerGCJobParam TriggerGCJob请求参数
type TriggerGCJobParam struct {
	BaseParam
	Params TriggerGCJobDetailParam `json:"params"` // 详细参数
}

