// Copyright (c) ZStack.io, Inc.

package param

// BatchAddBareMetal2IpmiChassisDetailParam BatchAddBareMetal2IpmiChassis详细参数
type BatchAddBareMetal2IpmiChassisDetailParam struct {
	rest string `json:"chassisInfo" validate:"required"` // 必填
	rest string `json:"longJobName,omitempty"`
	rest string `json:"longJobDescription,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// BatchAddBareMetal2IpmiChassisParam BatchAddBareMetal2IpmiChassis请求参数
type BatchAddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params BatchAddBareMetal2IpmiChassisDetailParam `json:"params"` // 详细参数
}

