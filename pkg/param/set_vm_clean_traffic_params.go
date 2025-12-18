// Copyright (c) ZStack.io, Inc.

package param

// SetVmCleanTrafficDetailParam SetVmCleanTraffic detail param
type SetVmCleanTrafficDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetVmCleanTrafficParam SetVmCleanTraffic request param
type SetVmCleanTrafficParam struct {
	BaseParam
	Params SetVmCleanTrafficDetailParam `json:"params"`
}
