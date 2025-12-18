// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunVirtualRouterDetailParam UpdateAliyunVirtualRouter detail param
type UpdateAliyunVirtualRouterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunVirtualRouterParam UpdateAliyunVirtualRouter request param
type UpdateAliyunVirtualRouterParam struct {
	BaseParam
	Params UpdateAliyunVirtualRouterDetailParam `json:"params"`
}
