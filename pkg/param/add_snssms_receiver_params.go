// Copyright (c) ZStack.io, Inc.

package param

// AddSNSSmsReceiverDetailParam AddSNSSmsReceiver详细参数
type AddSNSSmsReceiverDetailParam struct {
	rest string `json:"phoneNumber,omitempty"`
	rest []string `json:"phoneNumberList,omitempty"`
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSNSSmsReceiverParam AddSNSSmsReceiver请求参数
type AddSNSSmsReceiverParam struct {
	BaseParam
	Params AddSNSSmsReceiverDetailParam `json:"params"` // 详细参数
}

