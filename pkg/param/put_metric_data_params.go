// Copyright (c) ZStack.io, Inc.

package param

// PutMetricDataDetailParam PutMetricData详细参数
type PutMetricDataDetailParam struct {
	rest string `json:"namespace" validate:"required"` // 必填
	rest []interface{} `json:"data" validate:"required"` // 必填
}

// PutMetricDataParam PutMetricData请求参数
type PutMetricDataParam struct {
	BaseParam
	Params PutMetricDataDetailParam `json:"params"` // 详细参数
}

