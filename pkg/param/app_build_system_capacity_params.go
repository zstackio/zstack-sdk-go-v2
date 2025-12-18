// Copyright (c) ZStack.io, Inc.

package param

// GetAppBuildSystemCapacityDetailParam GetAppBuildSystemCapacity详细参数
type GetAppBuildSystemCapacityDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetAppBuildSystemCapacityParam GetAppBuildSystemCapacity请求参数
type GetAppBuildSystemCapacityParam struct {
	BaseParam
	Params GetAppBuildSystemCapacityDetailParam `json:"params"` // 详细参数
}

