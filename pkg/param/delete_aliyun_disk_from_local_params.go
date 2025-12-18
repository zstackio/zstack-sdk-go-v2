// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunDiskFromLocalDetailParam DeleteAliyunDiskFromLocal detail param
type DeleteAliyunDiskFromLocalDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunDiskFromLocalParam DeleteAliyunDiskFromLocal request param
type DeleteAliyunDiskFromLocalParam struct {
	BaseParam
	Params DeleteAliyunDiskFromLocalDetailParam `json:"params"`
}
