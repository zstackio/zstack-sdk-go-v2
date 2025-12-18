// Copyright (c) ZStack.io, Inc.

package param

// PullSdnControllerTenantDetailParam PullSdnControllerTenant详细参数
type PullSdnControllerTenantDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// PullSdnControllerTenantParam PullSdnControllerTenant请求参数
type PullSdnControllerTenantParam struct {
	BaseParam
	Params PullSdnControllerTenantDetailParam `json:"params"` // 详细参数
}

