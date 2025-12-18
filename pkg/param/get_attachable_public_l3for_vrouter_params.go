// Copyright (c) ZStack.io, Inc.

package param

// GetAttachablePublicL3ForVRouterDetailParam GetAttachablePublicL3ForVRouter detail param
type GetAttachablePublicL3ForVRouterDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetAttachablePublicL3ForVRouterParam GetAttachablePublicL3ForVRouter request param
type GetAttachablePublicL3ForVRouterParam struct {
	BaseParam
	Params GetAttachablePublicL3ForVRouterDetailParam `json:"params"`
}
