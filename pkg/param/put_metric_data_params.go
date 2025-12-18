// Copyright (c) ZStack.io, Inc.

package param

// PutMetricDataDetailParam PutMetricData detail param
type PutMetricDataDetailParam struct {
	Namespace string `json:"namespace" validate:"required"`
	Data []interface{} `json:"data" validate:"required"`
}

// PutMetricDataParam PutMetricData request param
type PutMetricDataParam struct {
	BaseParam
	Params PutMetricDataDetailParam `json:"params"`
}
