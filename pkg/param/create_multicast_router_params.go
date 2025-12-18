// Copyright (c) ZStack.io, Inc.

package param

// CreateMulticastRouterDetailParam CreateMulticastRouter detail param
type CreateMulticastRouterDetailParam struct {
	VpcRouterVmUuid string `json:"vpcRouterVmUuid" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMulticastRouterParam CreateMulticastRouter request param
type CreateMulticastRouterParam struct {
	BaseParam
	Params CreateMulticastRouterDetailParam `json:"params"`
}
