// Copyright (c) ZStack.io, Inc.

package param

// DeleteHuaweiIMasterTenantDetailParam DeleteHuaweiIMasterTenant detail param
type DeleteHuaweiIMasterTenantDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterTenantParam DeleteHuaweiIMasterTenant request param
type DeleteHuaweiIMasterTenantParam struct {
	BaseParam
	Params DeleteHuaweiIMasterTenantDetailParam `json:"params"`
}
