// Copyright (c) ZStack.io, Inc.

package param

// StopAllResourcesInIAM2ProjectDetailParam StopAllResourcesInIAM2Project detail param
type StopAllResourcesInIAM2ProjectDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// StopAllResourcesInIAM2ProjectParam StopAllResourcesInIAM2Project request param
type StopAllResourcesInIAM2ProjectParam struct {
	BaseParam
	Params StopAllResourcesInIAM2ProjectDetailParam `json:"params"`
}
