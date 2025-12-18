// Copyright (c) ZStack.io, Inc.

package param

// AttachAppBuildSystemToZoneDetailParam AttachAppBuildSystemToZone detail param
type AttachAppBuildSystemToZoneDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	BuildSystemUuid string `json:"buildSystemUuid" validate:"required"`
}

// AttachAppBuildSystemToZoneParam AttachAppBuildSystemToZone request param
type AttachAppBuildSystemToZoneParam struct {
	BaseParam
	Params AttachAppBuildSystemToZoneDetailParam `json:"params"`
}
