// Copyright (c) ZStack.io, Inc.

package param

// BatchCreateBaremetalChassisDetailParam BatchCreateBaremetalChassis详细参数
type BatchCreateBaremetalChassisDetailParam struct {
	rest string `json:"baremetalChassisInfo" validate:"required"` // 必填
	rest string `json:"longJobName,omitempty"`
	rest string `json:"longJobDescription,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// BatchCreateBaremetalChassisParam BatchCreateBaremetalChassis请求参数
type BatchCreateBaremetalChassisParam struct {
	BaseParam
	Params BatchCreateBaremetalChassisDetailParam `json:"params"` // 详细参数
}

