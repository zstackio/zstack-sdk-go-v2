// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunRouterInterfaceRemoteDetailParam CreateAliyunRouterInterfaceRemote详细参数
type CreateAliyunRouterInterfaceRemoteDetailParam struct {
	rest string `json:"dataCenterUuid" validate:"required"` // 必填
	rest string `json:"accessPointUuid,omitempty"`
	rest string `json:"spec,omitempty"`
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest string `json:"routerType" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunRouterInterfaceRemoteParam CreateAliyunRouterInterfaceRemote请求参数
type CreateAliyunRouterInterfaceRemoteParam struct {
	BaseParam
	Params CreateAliyunRouterInterfaceRemoteDetailParam `json:"params"` // 详细参数
}

