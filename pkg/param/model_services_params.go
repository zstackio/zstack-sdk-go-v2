// Copyright (c) ZStack.io, Inc.

package param

// DeleteModelServicesDetailParam DeleteModelServices详细参数
type DeleteModelServicesDetailParam struct {
	rest []string `json:"uuids" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteModelServicesParam DeleteModelServices请求参数
type DeleteModelServicesParam struct {
	BaseParam
	Params DeleteModelServicesDetailParam `json:"params"` // 详细参数
}

