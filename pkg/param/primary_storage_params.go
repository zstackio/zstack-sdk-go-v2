// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// ReconnectPrimaryStorageParamDetail ReconnectPrimaryStorage detail param
type ReconnectPrimaryStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectPrimaryStorageParam ReconnectPrimaryStorage request param
type ReconnectPrimaryStorageParam struct {
	BaseParam
	ReconnectPrimaryStorage ReconnectPrimaryStorageParamDetail `json:"reconnectPrimaryStorage"`
}
// UpdatePrimaryStorageParamDetail UpdatePrimaryStorage detail param
type UpdatePrimaryStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
}

// UpdatePrimaryStorageParam UpdatePrimaryStorage request param
type UpdatePrimaryStorageParam struct {
	BaseParam
	UpdatePrimaryStorage UpdatePrimaryStorageParamDetail `json:"updatePrimaryStorage"`
}
// DeletePrimaryStorageParamDetail DeletePrimaryStorage detail param
type DeletePrimaryStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePrimaryStorageParam DeletePrimaryStorage request param
type DeletePrimaryStorageParam struct {
	BaseParam
	DeletePrimaryStorage DeletePrimaryStorageParamDetail `json:"deletePrimaryStorage"`
}
