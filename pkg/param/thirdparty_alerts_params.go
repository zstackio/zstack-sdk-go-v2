// Copyright (c) ZStack.io, Inc.

package param

// UpdateThirdpartyAlertsDetailParam UpdateThirdpartyAlerts详细参数
type UpdateThirdpartyAlertsDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest int64 `json:"startTimeMillis,omitempty"`
	rest int64 `json:"endTimeMillis,omitempty"`
	rest string `json:"updateReadStatus,omitempty"`
}

// UpdateThirdpartyAlertsParam UpdateThirdpartyAlerts请求参数
type UpdateThirdpartyAlertsParam struct {
	BaseParam
	Params UpdateThirdpartyAlertsDetailParam `json:"params"` // 详细参数
}

