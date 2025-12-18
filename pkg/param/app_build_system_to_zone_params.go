// Copyright (c) ZStack.io, Inc.

package param

// DetachAppBuildSystemToZoneDetailParam DetachAppBuildSystemToZone详细参数
type DetachAppBuildSystemToZoneDetailParam struct {
	rest string `json:"buildSystemUuid" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
}

// DetachAppBuildSystemToZoneParam DetachAppBuildSystemToZone请求参数
type DetachAppBuildSystemToZoneParam struct {
	BaseParam
	Params DetachAppBuildSystemToZoneDetailParam `json:"params"` // 详细参数
}

