// Copyright (c) ZStack.io, Inc.

package param

// PowerOffHostDetailParam PowerOffHost detail param
type PowerOffHostDetailParam struct {
	AdminPassword string `json:"adminPassword" validate:"required"`
	HostUuids []string `json:"hostUuids" validate:"required"`
	WaitTaskCompleted bool `json:"waitTaskCompleted,omitempty"`
	MaxWaitTime int64 `json:"maxWaitTime,omitempty"`
}

// PowerOffHostParam PowerOffHost request param
type PowerOffHostParam struct {
	BaseParam
	Params PowerOffHostDetailParam `json:"params"`
}
