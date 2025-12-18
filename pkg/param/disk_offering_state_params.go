// Copyright (c) ZStack.io, Inc.

package param

// ChangeDiskOfferingStateDetailParam ChangeDiskOfferingState详细参数
type ChangeDiskOfferingStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeDiskOfferingStateParam ChangeDiskOfferingState请求参数
type ChangeDiskOfferingStateParam struct {
	BaseParam
	Params ChangeDiskOfferingStateDetailParam `json:"params"` // 详细参数
}

