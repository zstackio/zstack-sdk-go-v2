// Copyright (c) ZStack.io, Inc.

package param

// DeleteNasMountTargetDetailParam DeleteNasMountTarget detail param
type DeleteNasMountTargetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteNasMountTargetParam DeleteNasMountTarget request param
type DeleteNasMountTargetParam struct {
	BaseParam
	Params DeleteNasMountTargetDetailParam `json:"params"`
}
