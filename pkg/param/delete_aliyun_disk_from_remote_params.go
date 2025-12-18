// Copyright (c) ZStack.io, Inc.

package param

// DeleteAliyunDiskFromRemoteDetailParam DeleteAliyunDiskFromRemote detail param
type DeleteAliyunDiskFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAliyunDiskFromRemoteParam DeleteAliyunDiskFromRemote request param
type DeleteAliyunDiskFromRemoteParam struct {
	BaseParam
	Params DeleteAliyunDiskFromRemoteDetailParam `json:"params"`
}
