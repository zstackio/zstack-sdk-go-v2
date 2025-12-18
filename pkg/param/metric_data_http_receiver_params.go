// Copyright (c) ZStack.io, Inc.

package param

// CreateMetricDataHttpReceiverDetailParam CreateMetricDataHttpReceiver详细参数
type CreateMetricDataHttpReceiverDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest bool `json:"defaultEnable,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateMetricDataHttpReceiverParam CreateMetricDataHttpReceiver请求参数
type CreateMetricDataHttpReceiverParam struct {
	BaseParam
	Params CreateMetricDataHttpReceiverDetailParam `json:"params"` // 详细参数
}

