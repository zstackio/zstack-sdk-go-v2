// Copyright (c) ZStack.io, Inc.

package param

// RemoveMonFromCephPrimaryStorageDetailParam RemoveMonFromCephPrimaryStorage detail param
type RemoveMonFromCephPrimaryStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	MonHostnames []string `json:"monHostnames" validate:"required"`
}

// RemoveMonFromCephPrimaryStorageParam RemoveMonFromCephPrimaryStorage request param
type RemoveMonFromCephPrimaryStorageParam struct {
	BaseParam
	Params RemoveMonFromCephPrimaryStorageDetailParam `json:"params"`
}
