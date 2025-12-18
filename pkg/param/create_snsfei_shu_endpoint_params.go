// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSFeiShuEndpointDetailParam CreateSNSFeiShuEndpoint detail param
type CreateSNSFeiShuEndpointDetailParam struct {
	Url string `json:"url" validate:"required"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	Secret string `json:"secret,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSFeiShuEndpointParam CreateSNSFeiShuEndpoint request param
type CreateSNSFeiShuEndpointParam struct {
	BaseParam
	Params CreateSNSFeiShuEndpointDetailParam `json:"params"`
}
