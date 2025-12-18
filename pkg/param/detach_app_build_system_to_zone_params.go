// Copyright (c) ZStack.io, Inc.

package param

// DetachAppBuildSystemToZoneDetailParam DetachAppBuildSystemToZone detail param
type DetachAppBuildSystemToZoneDetailParam struct {
	BuildSystemUuid string `json:"buildSystemUuid" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
}

// DetachAppBuildSystemToZoneParam DetachAppBuildSystemToZone request param
type DetachAppBuildSystemToZoneParam struct {
	BaseParam
	Params DetachAppBuildSystemToZoneDetailParam `json:"params"`
}
