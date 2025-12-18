// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSPluginEndpointDetailParam CreateSNSPluginEndpoint detail param
type CreateSNSPluginEndpointDetailParam struct {
	Type string `json:"type" validate:"required"`
	TimeoutInSeconds int64 `json:"timeoutInSeconds" validate:"required"`
	Properties map[string]string `json:"properties,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSPluginEndpointParam CreateSNSPluginEndpoint request param
type CreateSNSPluginEndpointParam struct {
	BaseParam
	Params CreateSNSPluginEndpointDetailParam `json:"params"`
}
