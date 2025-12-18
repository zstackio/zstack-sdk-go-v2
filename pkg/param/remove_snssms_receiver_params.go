// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSSmsReceiverDetailParam RemoveSNSSmsReceiver详细参数
type RemoveSNSSmsReceiverDetailParam struct {
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"phoneNumber,omitempty"`
	rest []string `json:"phoneNumberList,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// RemoveSNSSmsReceiverParam RemoveSNSSmsReceiver请求参数
type RemoveSNSSmsReceiverParam struct {
	BaseParam
	Params RemoveSNSSmsReceiverDetailParam `json:"params"` // 详细参数
}

