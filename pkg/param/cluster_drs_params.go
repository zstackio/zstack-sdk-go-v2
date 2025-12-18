// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterDRSDetailParam UpdateClusterDRS详细参数
type UpdateClusterDRSDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"automationLevel,omitempty"`
	rest []interface{} `json:"thresholds,omitempty"`
	rest int `json:"thresholdDuration,omitempty"`
	rest string `json:"state,omitempty"`
}

// UpdateClusterDRSParam UpdateClusterDRS请求参数
type UpdateClusterDRSParam struct {
	BaseParam
	Params UpdateClusterDRSDetailParam `json:"params"` // 详细参数
}

