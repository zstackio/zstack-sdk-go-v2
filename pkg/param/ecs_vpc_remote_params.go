// Copyright (c) ZStack.io, Inc.

package param

// CreateEcsVpcRemoteDetailParam CreateEcsVpcRemote详细参数
type CreateEcsVpcRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"cidrBlock" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"vRouterName" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateEcsVpcRemoteParam CreateEcsVpcRemote请求参数
type CreateEcsVpcRemoteParam struct {
	BaseParam
	Params CreateEcsVpcRemoteDetailParam `json:"params"` // 详细参数
}

