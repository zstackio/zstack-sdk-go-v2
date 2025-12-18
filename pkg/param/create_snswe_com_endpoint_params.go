// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSWeComEndpointDetailParam CreateSNSWeComEndpoint detail param
type CreateSNSWeComEndpointDetailParam struct {
	Url string `json:"url" validate:"required"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSWeComEndpointParam CreateSNSWeComEndpoint request param
type CreateSNSWeComEndpointParam struct {
	BaseParam
	Params CreateSNSWeComEndpointDetailParam `json:"params"`
}
