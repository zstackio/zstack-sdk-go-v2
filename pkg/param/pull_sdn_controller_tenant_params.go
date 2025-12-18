// Copyright (c) ZStack.io, Inc.

package param

// PullSdnControllerTenantDetailParam PullSdnControllerTenant detail param
type PullSdnControllerTenantDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// PullSdnControllerTenantParam PullSdnControllerTenant request param
type PullSdnControllerTenantParam struct {
	BaseParam
	Params PullSdnControllerTenantDetailParam `json:"params"`
}
