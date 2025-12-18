// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunRouterInterfaceRemoteDetailParam CreateAliyunRouterInterfaceRemote detail param
type CreateAliyunRouterInterfaceRemoteDetailParam struct {
	DataCenterUuid string `json:"dataCenterUuid" validate:"required"`
	AccessPointUuid string `json:"accessPointUuid,omitempty"`
	Spec string `json:"spec,omitempty"`
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	RouterType string `json:"routerType" validate:"required"`
	Description string `json:"description,omitempty"`
	Name string `json:"name" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunRouterInterfaceRemoteParam CreateAliyunRouterInterfaceRemote request param
type CreateAliyunRouterInterfaceRemoteParam struct {
	BaseParam
	Params CreateAliyunRouterInterfaceRemoteDetailParam `json:"params"`
}
