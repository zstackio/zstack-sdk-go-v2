// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunVpcVirtualRouterEntryRemoteDetailParam CreateAliyunVpcVirtualRouterEntryRemote detail param
type CreateAliyunVpcVirtualRouterEntryRemoteDetailParam struct {
	VRouterUuid string `json:"vRouterUuid" validate:"required"`
	DstCidrBlock string `json:"dstCidrBlock" validate:"required"`
	NextHopUuid string `json:"nextHopUuid" validate:"required"`
	NextHopType string `json:"nextHopType" validate:"required"`
	VRouterType string `json:"vRouterType" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunVpcVirtualRouterEntryRemoteParam CreateAliyunVpcVirtualRouterEntryRemote request param
type CreateAliyunVpcVirtualRouterEntryRemoteParam struct {
	BaseParam
	Params CreateAliyunVpcVirtualRouterEntryRemoteDetailParam `json:"params"`
}
