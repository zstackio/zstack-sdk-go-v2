// Copyright (c) ZStack.io, Inc.

package param

// GetInterdependentL3NetworksImagesDetailParam GetInterdependentL3NetworksImages detail param
type GetInterdependentL3NetworksImagesDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	RaiseException bool `json:"raiseException,omitempty"`
}

// GetInterdependentL3NetworksImagesParam GetInterdependentL3NetworksImages request param
type GetInterdependentL3NetworksImagesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksImagesDetailParam `json:"params"`
}
