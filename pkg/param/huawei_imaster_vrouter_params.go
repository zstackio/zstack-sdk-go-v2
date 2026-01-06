// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// CreateHuaweiIMasterVRouterParamDetail CreateHuaweiIMasterVRouter detail param
type CreateHuaweiIMasterVRouterParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	HuaweiVpcUuid string `json:"huaweiVpcUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHuaweiIMasterVRouterParam CreateHuaweiIMasterVRouter request param
type CreateHuaweiIMasterVRouterParam struct {
	BaseParam
	Params CreateHuaweiIMasterVRouterParamDetail `json:"params"`
}
// DeleteHuaweiIMasterVRouterParamDetail DeleteHuaweiIMasterVRouter detail param
type DeleteHuaweiIMasterVRouterParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteHuaweiIMasterVRouterParam DeleteHuaweiIMasterVRouter request param
type DeleteHuaweiIMasterVRouterParam struct {
	BaseParam
	Params DeleteHuaweiIMasterVRouterParamDetail `json:"params"`
}
