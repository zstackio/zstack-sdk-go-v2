// Copyright (c) ZStack.io, Inc.

package param

// GetAllEventMetadataDetailParam GetAllEventMetadata detail param
type GetAllEventMetadataDetailParam struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// GetAllEventMetadataParam GetAllEventMetadata request param
type GetAllEventMetadataParam struct {
	BaseParam
	Params GetAllEventMetadataDetailParam `json:"params"`
}
