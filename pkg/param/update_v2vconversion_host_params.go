// Copyright (c) ZStack.io, Inc.

package param

// UpdateV2VConversionHostDetailParam UpdateV2VConversionHost detail param
type UpdateV2VConversionHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StoragePath string `json:"storagePath,omitempty"`
}

// UpdateV2VConversionHostParam UpdateV2VConversionHost request param
type UpdateV2VConversionHostParam struct {
	BaseParam
	Params UpdateV2VConversionHostDetailParam `json:"params"`
}
