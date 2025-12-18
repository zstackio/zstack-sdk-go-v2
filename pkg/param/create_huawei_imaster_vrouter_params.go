// Copyright (c) ZStack.io, Inc.

package param

// CreateHuaweiIMasterVRouterDetailParam CreateHuaweiIMasterVRouter detail param
type CreateHuaweiIMasterVRouterDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	HuaweiVpcUuid string `json:"huaweiVpcUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateHuaweiIMasterVRouterParam CreateHuaweiIMasterVRouter request param
type CreateHuaweiIMasterVRouterParam struct {
	BaseParam
	Params CreateHuaweiIMasterVRouterDetailParam `json:"params"`
}
