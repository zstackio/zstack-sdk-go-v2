// Copyright (c) ZStack.io, Inc.

package param

// AttachOssBucketToEcsDataCenterDetailParam AttachOssBucketToEcsDataCenter详细参数
type AttachOssBucketToEcsDataCenterDetailParam struct {
	rest string `json:"ossBucketUuid" validate:"required"` // 必填
}

// AttachOssBucketToEcsDataCenterParam AttachOssBucketToEcsDataCenter请求参数
type AttachOssBucketToEcsDataCenterParam struct {
	BaseParam
	Params AttachOssBucketToEcsDataCenterDetailParam `json:"params"` // 详细参数
}

