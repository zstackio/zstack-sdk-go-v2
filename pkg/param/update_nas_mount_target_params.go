// Copyright (c) ZStack.io, Inc.

package param

// UpdateNasMountTargetDetailParam UpdateNasMountTarget detail param
type UpdateNasMountTargetDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNasMountTargetParam UpdateNasMountTarget request param
type UpdateNasMountTargetParam struct {
	BaseParam
	Params UpdateNasMountTargetDetailParam `json:"params"`
}
