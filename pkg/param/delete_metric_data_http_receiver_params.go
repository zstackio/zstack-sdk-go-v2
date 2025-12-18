// Copyright (c) ZStack.io, Inc.

package param

// DeleteMetricDataHttpReceiverDetailParam DeleteMetricDataHttpReceiver detail param
type DeleteMetricDataHttpReceiverDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteMetricDataHttpReceiverParam DeleteMetricDataHttpReceiver request param
type DeleteMetricDataHttpReceiverParam struct {
	BaseParam
	Params DeleteMetricDataHttpReceiverDetailParam `json:"params"`
}
