// Copyright (c) ZStack.io, Inc.

package param

// ProvisionVirtualRouterConfigDetailParam ProvisionVirtualRouterConfig detail param
type ProvisionVirtualRouterConfigDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// ProvisionVirtualRouterConfigParam ProvisionVirtualRouterConfig request param
type ProvisionVirtualRouterConfigParam struct {
	BaseParam
	Params ProvisionVirtualRouterConfigDetailParam `json:"params"`
}
