// Copyright (c) ZStack.io, Inc.

package param

// DetachOssBucketFromEcsDataCenterDetailParam DetachOssBucketFromEcsDataCenter detail param
type DetachOssBucketFromEcsDataCenterDetailParam struct {
	OssBucketUuid string `json:"ossBucketUuid" validate:"required"`
}

// DetachOssBucketFromEcsDataCenterParam DetachOssBucketFromEcsDataCenter request param
type DetachOssBucketFromEcsDataCenterParam struct {
	BaseParam
	Params DetachOssBucketFromEcsDataCenterDetailParam `json:"params"`
}
