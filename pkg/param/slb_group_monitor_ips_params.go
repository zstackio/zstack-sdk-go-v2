// Copyright (c) ZStack.io, Inc.

package param

// ChangeSlbGroupMonitorIpsDetailParam ChangeSlbGroupMonitorIps详细参数
type ChangeSlbGroupMonitorIpsDetailParam struct {
	rest string `json:"slbGroupUuid" validate:"required"` // 必填
	rest []string `json:"monitorIps" validate:"required"` // 必填
}

// ChangeSlbGroupMonitorIpsParam ChangeSlbGroupMonitorIps请求参数
type ChangeSlbGroupMonitorIpsParam struct {
	BaseParam
	Params ChangeSlbGroupMonitorIpsDetailParam `json:"params"` // 详细参数
}

