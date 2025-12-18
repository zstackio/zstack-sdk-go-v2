// Copyright (c) ZStack.io, Inc.

package param

// GetAttachablePublicL3ForVRouterDetailParam GetAttachablePublicL3ForVRouter详细参数
type GetAttachablePublicL3ForVRouterDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetAttachablePublicL3ForVRouterParam GetAttachablePublicL3ForVRouter请求参数
type GetAttachablePublicL3ForVRouterParam struct {
	BaseParam
	Params GetAttachablePublicL3ForVRouterDetailParam `json:"params"` // 详细参数
}

