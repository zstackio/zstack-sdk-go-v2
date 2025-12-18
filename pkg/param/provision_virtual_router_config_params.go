// Copyright (c) ZStack.io, Inc.

package param

// ProvisionVirtualRouterConfigDetailParam ProvisionVirtualRouterConfig详细参数
type ProvisionVirtualRouterConfigDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// ProvisionVirtualRouterConfigParam ProvisionVirtualRouterConfig请求参数
type ProvisionVirtualRouterConfigParam struct {
	BaseParam
	Params ProvisionVirtualRouterConfigDetailParam `json:"params"` // 详细参数
}

