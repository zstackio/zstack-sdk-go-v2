// Copyright (c) ZStack.io, Inc.

package param

// DetachOssBucketFromEcsDataCenterDetailParam DetachOssBucketFromEcsDataCenter详细参数
type DetachOssBucketFromEcsDataCenterDetailParam struct {
	rest string `json:"ossBucketUuid" validate:"required"` // 必填
}

// DetachOssBucketFromEcsDataCenterParam DetachOssBucketFromEcsDataCenter请求参数
type DetachOssBucketFromEcsDataCenterParam struct {
	BaseParam
	Params DetachOssBucketFromEcsDataCenterDetailParam `json:"params"` // 详细参数
}

