// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateNasMountTargetParamDetail UpdateNasMountTarget detail param
type UpdateNasMountTargetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNasMountTargetParam UpdateNasMountTarget request param
type UpdateNasMountTargetParam struct {
	BaseParam
	UpdateNasMountTarget UpdateNasMountTargetParamDetail `json:"updateNasMountTarget"`
}
// DeleteNasMountTargetParamDetail DeleteNasMountTarget detail param
type DeleteNasMountTargetParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNasMountTargetParam DeleteNasMountTarget request param
type DeleteNasMountTargetParam struct {
	BaseParam
	DeleteNasMountTarget DeleteNasMountTargetParamDetail `json:"deleteNasMountTarget"`
}
