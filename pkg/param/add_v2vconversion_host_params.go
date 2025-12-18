// Copyright (c) ZStack.io, Inc.

package param

// AddV2VConversionHostDetailParam AddV2VConversionHost detail param
type AddV2VConversionHostDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	StoragePath string `json:"storagePath" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddV2VConversionHostParam AddV2VConversionHost request param
type AddV2VConversionHostParam struct {
	BaseParam
	Params AddV2VConversionHostDetailParam `json:"params"`
}
