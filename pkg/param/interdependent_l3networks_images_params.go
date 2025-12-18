// Copyright (c) ZStack.io, Inc.

package param

// GetInterdependentL3NetworksImagesDetailParam GetInterdependentL3NetworksImages详细参数
type GetInterdependentL3NetworksImagesDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest bool `json:"raiseException,omitempty"`
}

// GetInterdependentL3NetworksImagesParam GetInterdependentL3NetworksImages请求参数
type GetInterdependentL3NetworksImagesParam struct {
	BaseParam
	Params GetInterdependentL3NetworksImagesDetailParam `json:"params"` // 详细参数
}

