// Copyright (c) ZStack.io, Inc.

package param

// AddBareMetal2IpmiChassisDetailParam AddBareMetal2IpmiChassis详细参数
type AddBareMetal2IpmiChassisDetailParam struct {
	rest string `json:"ipmiAddress" validate:"required"` // 必填
	rest int `json:"ipmiPort,omitempty"`
	rest string `json:"ipmiUsername" validate:"required"` // 必填
	rest string `json:"ipmiPassword" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"provisionType,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddBareMetal2IpmiChassisParam AddBareMetal2IpmiChassis请求参数
type AddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params AddBareMetal2IpmiChassisDetailParam `json:"params"` // 详细参数
}

