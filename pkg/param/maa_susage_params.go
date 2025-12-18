// Copyright (c) ZStack.io, Inc.

package param

// GetMaaSUsageDetailParam GetMaaSUsage详细参数
type GetMaaSUsageDetailParam struct {
}

// GetMaaSUsageParam GetMaaSUsage请求参数
type GetMaaSUsageParam struct {
	BaseParam
	Params GetMaaSUsageDetailParam `json:"params"` // 详细参数
}

