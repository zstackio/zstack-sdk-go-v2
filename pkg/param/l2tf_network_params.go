// Copyright (c) ZStack.io, Inc.

package param

// CreateL2TfNetworkDetailParam CreateL2TfNetwork详细参数
type CreateL2TfNetworkDetailParam struct {
	rest string `json:"ipPrefix,omitempty"`
	rest int `json:"ipPrefixLength,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"` // 必填
	rest string `json:"physicalInterface,omitempty"` // 必填
	rest string `json:"type,omitempty"`
	rest string `json:"vSwitchType,omitempty"`
	rest bool `json:"isolated,omitempty"`
	rest string `json:"pvlan,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateL2TfNetworkParam CreateL2TfNetwork请求参数
type CreateL2TfNetworkParam struct {
	BaseParam
	Params CreateL2TfNetworkDetailParam `json:"params"` // 详细参数
}

