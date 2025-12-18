// Copyright (c) ZStack.io, Inc.

package param

// CreateVRouterRouteTableDetailParam CreateVRouterRouteTable detail param
type CreateVRouterRouteTableDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVRouterRouteTableParam CreateVRouterRouteTable request param
type CreateVRouterRouteTableParam struct {
	BaseParam
	Params CreateVRouterRouteTableDetailParam `json:"params"`
}
