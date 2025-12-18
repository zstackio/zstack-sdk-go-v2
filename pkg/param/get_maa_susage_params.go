// Copyright (c) ZStack.io, Inc.

package param

// GetMaaSUsageDetailParam GetMaaSUsage detail param
type GetMaaSUsageDetailParam struct {
}

// GetMaaSUsageParam GetMaaSUsage request param
type GetMaaSUsageParam struct {
	BaseParam
	Params GetMaaSUsageDetailParam `json:"params"`
}
