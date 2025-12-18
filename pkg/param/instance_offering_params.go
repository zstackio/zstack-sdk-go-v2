// Copyright (c) ZStack.io, Inc.

package param

// DeleteInstanceOfferingDetailParam DeleteInstanceOffering详细参数
type DeleteInstanceOfferingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteInstanceOfferingParam DeleteInstanceOffering请求参数
type DeleteInstanceOfferingParam struct {
	BaseParam
	Params DeleteInstanceOfferingDetailParam `json:"params"` // 详细参数
}

