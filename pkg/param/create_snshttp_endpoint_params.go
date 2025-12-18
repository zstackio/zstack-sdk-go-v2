// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSHttpEndpointDetailParam CreateSNSHttpEndpoint detail param
type CreateSNSHttpEndpointDetailParam struct {
	Url string `json:"url" validate:"required"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSHttpEndpointParam CreateSNSHttpEndpoint request param
type CreateSNSHttpEndpointParam struct {
	BaseParam
	Params CreateSNSHttpEndpointDetailParam `json:"params"`
}
