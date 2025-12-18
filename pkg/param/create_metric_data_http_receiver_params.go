// Copyright (c) ZStack.io, Inc.

package param

// CreateMetricDataHttpReceiverDetailParam CreateMetricDataHttpReceiver detail param
type CreateMetricDataHttpReceiverDetailParam struct {
	Name string `json:"name" validate:"required"`
	Url string `json:"url" validate:"required"`
	Description string `json:"description,omitempty"`
	DefaultEnable bool `json:"defaultEnable,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMetricDataHttpReceiverParam CreateMetricDataHttpReceiver request param
type CreateMetricDataHttpReceiverParam struct {
	BaseParam
	Params CreateMetricDataHttpReceiverDetailParam `json:"params"`
}
