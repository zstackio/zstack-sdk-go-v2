// Copyright (c) ZStack.io, Inc.

package param

// BatchAddBareMetal2IpmiChassisDetailParam BatchAddBareMetal2IpmiChassis detail param
type BatchAddBareMetal2IpmiChassisDetailParam struct {
	ChassisInfo string `json:"chassisInfo" validate:"required"`
	LongJobName string `json:"longJobName,omitempty"`
	LongJobDescription string `json:"longJobDescription,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchAddBareMetal2IpmiChassisParam BatchAddBareMetal2IpmiChassis request param
type BatchAddBareMetal2IpmiChassisParam struct {
	BaseParam
	Params BatchAddBareMetal2IpmiChassisDetailParam `json:"params"`
}
