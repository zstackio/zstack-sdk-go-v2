// Copyright (c) ZStack.io, Inc.

package param

// CreateHuaweiIMasterVRouterDetailParam CreateHuaweiIMasterVRouter详细参数
type CreateHuaweiIMasterVRouterDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"huaweiVpcUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateHuaweiIMasterVRouterParam CreateHuaweiIMasterVRouter请求参数
type CreateHuaweiIMasterVRouterParam struct {
	BaseParam
	Params CreateHuaweiIMasterVRouterDetailParam `json:"params"` // 详细参数
}

