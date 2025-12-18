// Copyright (c) ZStack.io, Inc.

package param

// ChangeSNSApplicationPlatformStateDetailParam ChangeSNSApplicationPlatformState detail param
type ChangeSNSApplicationPlatformStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSNSApplicationPlatformStateParam ChangeSNSApplicationPlatformState request param
type ChangeSNSApplicationPlatformStateParam struct {
	BaseParam
	Params ChangeSNSApplicationPlatformStateDetailParam `json:"params"`
}
