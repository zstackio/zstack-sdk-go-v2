// Copyright (c) ZStack.io, Inc.

package param

// SetVmCleanTrafficDetailParam SetVmCleanTraffic详细参数
type SetVmCleanTrafficDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"enable" validate:"required"` // 必填
}

// SetVmCleanTrafficParam SetVmCleanTraffic请求参数
type SetVmCleanTrafficParam struct {
	BaseParam
	Params SetVmCleanTrafficDetailParam `json:"params"` // 详细参数
}

