// Copyright (c) ZStack.io, Inc.

package param

// AttachOssBucketToEcsDataCenterDetailParam AttachOssBucketToEcsDataCenter detail param
type AttachOssBucketToEcsDataCenterDetailParam struct {
	OssBucketUuid string `json:"ossBucketUuid" validate:"required"`
}

// AttachOssBucketToEcsDataCenterParam AttachOssBucketToEcsDataCenter request param
type AttachOssBucketToEcsDataCenterParam struct {
	BaseParam
	Params AttachOssBucketToEcsDataCenterDetailParam `json:"params"`
}
