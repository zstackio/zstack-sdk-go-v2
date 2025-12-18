// Copyright (c) ZStack.io, Inc.

package param

// AddModelCenterDetailParam AddModelCenter详细参数
type AddModelCenterDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest int `json:"managementPort" validate:"required"` // 必填
	rest string `json:"parameters,omitempty"`
	rest string `json:"storageNetworkUuid,omitempty"`
	rest string `json:"serviceNetworkUuid,omitempty"`
	rest string `json:"containerRegistry,omitempty"`
	rest string `json:"containerNetwork,omitempty"`
	rest string `json:"containerStorageNetwork,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddModelCenterParam AddModelCenter请求参数
type AddModelCenterParam struct {
	BaseParam
	Params AddModelCenterDetailParam `json:"params"` // 详细参数
}

