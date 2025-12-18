// Copyright (c) ZStack.io, Inc.

package param

// PowerOffHostDetailParam PowerOffHost详细参数
type PowerOffHostDetailParam struct {
	rest string `json:"adminPassword" validate:"required"` // 必填
	rest []string `json:"hostUuids" validate:"required"` // 必填
	rest bool `json:"waitTaskCompleted,omitempty"`
	rest int64 `json:"maxWaitTime,omitempty"`
}

// PowerOffHostParam PowerOffHost请求参数
type PowerOffHostParam struct {
	BaseParam
	Params PowerOffHostDetailParam `json:"params"` // 详细参数
}

