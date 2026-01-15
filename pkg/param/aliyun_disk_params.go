// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateAliyunDiskParamDetail UpdateAliyunDisk detail param
type UpdateAliyunDiskParamDetail struct {
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
	Params UpdateAliyunDiskParamDetail `json:"updateAliyunDisk"`
}
