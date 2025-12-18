// Copyright (c) ZStack.io, Inc.

package param

// ChangeVpcHaGroupMonitorIpsDetailParam ChangeVpcHaGroupMonitorIps详细参数
type ChangeVpcHaGroupMonitorIpsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []string `json:"monitorIps,omitempty"`
}

// ChangeVpcHaGroupMonitorIpsParam ChangeVpcHaGroupMonitorIps请求参数
type ChangeVpcHaGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeVpcHaGroupMonitorIpsDetailParam `json:"params"` // 详细参数
}

