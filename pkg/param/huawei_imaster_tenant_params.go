// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteHuaweiIMasterTenantParamDetail DeleteHuaweiIMasterTenant detail param
type DeleteHuaweiIMasterTenantParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterTenantParam DeleteHuaweiIMasterTenant request param
type DeleteHuaweiIMasterTenantParam struct {
	BaseParam
	Params DeleteHuaweiIMasterTenantParamDetail `json:"deleteHuaweiIMasterTenant"`
}
