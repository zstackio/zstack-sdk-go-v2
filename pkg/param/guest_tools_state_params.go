// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateGuestToolsStateParamDetail UpdateGuestToolsState detail param
type UpdateGuestToolsStateParamDetail struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// UpdateGuestToolsStateParam UpdateGuestToolsState request param
type UpdateGuestToolsStateParam struct {
	BaseParam
	UpdateGuestToolsState UpdateGuestToolsStateParamDetail `json:"updateGuestToolsState"`
}
