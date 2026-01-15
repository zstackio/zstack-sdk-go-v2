// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteVRouterRouteTableParamDetail DeleteVRouterRouteTable detail param
type DeleteVRouterRouteTableParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVRouterRouteTableParam DeleteVRouterRouteTable request param
type DeleteVRouterRouteTableParam struct {
	BaseParam
	Params DeleteVRouterRouteTableParamDetail `json:"deleteVRouterRouteTable"`
}
// GetVRouterRouteTableParamDetail GetVRouterRouteTable detail param
type GetVRouterRouteTableParamDetail struct {
	VirtualRouterVmUuid string `json:"virtualRouterVmUuid" validate:"required"`
}

// GetVRouterRouteTableParam GetVRouterRouteTable request param
type GetVRouterRouteTableParam struct {
	BaseParam
	Params GetVRouterRouteTableParamDetail `json:"getVRouterRouteTable"`
}
// CreateVRouterRouteTableParamDetail CreateVRouterRouteTable detail param
type CreateVRouterRouteTableParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVRouterRouteTableParam CreateVRouterRouteTable request param
type CreateVRouterRouteTableParam struct {
	BaseParam
	Params CreateVRouterRouteTableParamDetail `json:"createVRouterRouteTable"`
}
// UpdateVRouterRouteTableParamDetail UpdateVRouterRouteTable detail param
type UpdateVRouterRouteTableParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVRouterRouteTableParam UpdateVRouterRouteTable request param
type UpdateVRouterRouteTableParam struct {
	BaseParam
	Params UpdateVRouterRouteTableParamDetail `json:"updateVRouterRouteTable"`
}
