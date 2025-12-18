// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterTenantDetailParam DeleteHuaweiIMasterTenant详细参数
type DeleteHuaweiIMasterTenantDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"sdnControllerUuid,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterTenantParam DeleteHuaweiIMasterTenant请求参数
type DeleteHuaweiIMasterTenantParam struct {
	BaseParam
	Params DeleteHuaweiIMasterTenantDetailParam `json:"params"` // 详细参数
}

