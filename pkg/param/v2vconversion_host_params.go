// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddV2VConversionHostParamDetail AddV2VConversionHost detail param
type AddV2VConversionHostParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	StoragePath string `json:"storagePath" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddV2VConversionHostParam AddV2VConversionHost request param
type AddV2VConversionHostParam struct {
	BaseParam
	Params AddV2VConversionHostParamDetail `json:"params"`
}
// UpdateV2VConversionHostParamDetail UpdateV2VConversionHost detail param
type UpdateV2VConversionHostParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	StoragePath *string `json:"storagePath,omitempty"`
}

// UpdateV2VConversionHostParam UpdateV2VConversionHost request param
type UpdateV2VConversionHostParam struct {
	BaseParam
	Params UpdateV2VConversionHostParamDetail `json:"updateV2VConversionHost"`
}
// DeleteV2VConversionHostParamDetail DeleteV2VConversionHost detail param
type DeleteV2VConversionHostParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteV2VConversionHostParam DeleteV2VConversionHost request param
type DeleteV2VConversionHostParam struct {
	BaseParam
	Params DeleteV2VConversionHostParamDetail `json:"deleteV2VConversionHost"`
}
