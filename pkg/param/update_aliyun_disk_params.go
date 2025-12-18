// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunDiskDetailParam UpdateAliyunDisk detail param
type UpdateAliyunDiskDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DeleteWithInstance bool `json:"deleteWithInstance,omitempty"`
	DeleteAutoSnapshot bool `json:"deleteAutoSnapshot,omitempty"`
	EnableAutoSnapshot bool `json:"enableAutoSnapshot,omitempty"`
}

// UpdateAliyunDiskParam UpdateAliyunDisk request param
type UpdateAliyunDiskParam struct {
	BaseParam
	Params UpdateAliyunDiskDetailParam `json:"params"`
}
