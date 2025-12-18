// Copyright (c) ZStack.io, Inc.

package param

// GetAppBuildSystemCapacityDetailParam GetAppBuildSystemCapacity detail param
type GetAppBuildSystemCapacityDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetAppBuildSystemCapacityParam GetAppBuildSystemCapacity request param
type GetAppBuildSystemCapacityParam struct {
	BaseParam
	Params GetAppBuildSystemCapacityDetailParam `json:"params"`
}
