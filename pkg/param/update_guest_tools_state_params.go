// Copyright (c) ZStack.io, Inc.

package param

// UpdateGuestToolsStateDetailParam UpdateGuestToolsState detail param
type UpdateGuestToolsStateDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// UpdateGuestToolsStateParam UpdateGuestToolsState request param
type UpdateGuestToolsStateParam struct {
	BaseParam
	Params UpdateGuestToolsStateDetailParam `json:"params"`
}
