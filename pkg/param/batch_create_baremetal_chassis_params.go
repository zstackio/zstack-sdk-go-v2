// Copyright (c) ZStack.io, Inc.

package param

// BatchCreateBaremetalChassisDetailParam BatchCreateBaremetalChassis detail param
type BatchCreateBaremetalChassisDetailParam struct {
	BaremetalChassisInfo string `json:"baremetalChassisInfo" validate:"required"`
	LongJobName string `json:"longJobName,omitempty"`
	LongJobDescription string `json:"longJobDescription,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// BatchCreateBaremetalChassisParam BatchCreateBaremetalChassis request param
type BatchCreateBaremetalChassisParam struct {
	BaseParam
	Params BatchCreateBaremetalChassisDetailParam `json:"params"`
}
