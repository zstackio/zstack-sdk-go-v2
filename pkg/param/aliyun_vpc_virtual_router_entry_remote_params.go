// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunVpcVirtualRouterEntryRemoteDetailParam CreateAliyunVpcVirtualRouterEntryRemote详细参数
type CreateAliyunVpcVirtualRouterEntryRemoteDetailParam struct {
	rest string `json:"vRouterUuid" validate:"required"` // 必填
	rest string `json:"dstCidrBlock" validate:"required"` // 必填
	rest string `json:"nextHopUuid" validate:"required"` // 必填
	rest string `json:"nextHopType" validate:"required"` // 必填
	rest string `json:"vRouterType" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateAliyunVpcVirtualRouterEntryRemoteParam CreateAliyunVpcVirtualRouterEntryRemote请求参数
type CreateAliyunVpcVirtualRouterEntryRemoteParam struct {
	BaseParam
	Params CreateAliyunVpcVirtualRouterEntryRemoteDetailParam `json:"params"` // 详细参数
}

