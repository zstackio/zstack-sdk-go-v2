// Copyright (c) ZStack.io, Inc.

package param

// GetModelCenterServicesDetailParam GetModelCenterServices详细参数
type GetModelCenterServicesDetailParam struct {
	rest []string `json:"modelCenterUuids,omitempty"`
}

// GetModelCenterServicesParam GetModelCenterServices请求参数
type GetModelCenterServicesParam struct {
	BaseParam
	Params GetModelCenterServicesDetailParam `json:"params"` // 详细参数
}

