// Copyright (c) ZStack.io, Inc.

package param

// AddVCenterDetailParam AddVCenter详细参数
type AddVCenterDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest bool `json:"https,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"domainName" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddVCenterParam AddVCenter请求参数
type AddVCenterParam struct {
	BaseParam
	Params AddVCenterDetailParam `json:"params"` // 详细参数
}

