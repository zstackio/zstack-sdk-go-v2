// Copyright (c) ZStack.io, Inc.

package param

// AddVRouterRouteEntryDetailParam AddVRouterRouteEntry detail param
type AddVRouterRouteEntryDetailParam struct {
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	RouteTableUuid string `json:"routeTableUuid" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Target string `json:"target,omitempty"`
	Distance int `json:"distance,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddVRouterRouteEntryParam AddVRouterRouteEntry request param
type AddVRouterRouteEntryParam struct {
	BaseParam
	Params AddVRouterRouteEntryDetailParam `json:"params"`
}
