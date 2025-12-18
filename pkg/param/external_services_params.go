// Copyright (c) ZStack.io, Inc.

package param

// GetExternalServicesDetailParam GetExternalServices详细参数
type GetExternalServicesDetailParam struct {
}

// GetExternalServicesParam GetExternalServices请求参数
type GetExternalServicesParam struct {
	BaseParam
	Params GetExternalServicesDetailParam `json:"params"` // 详细参数
}

