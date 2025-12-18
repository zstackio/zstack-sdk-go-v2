// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSDingTalkEndpointDetailParam CreateSNSDingTalkEndpoint detail param
type CreateSNSDingTalkEndpointDetailParam struct {
	Url string `json:"url" validate:"required"`
	AtAll bool `json:"atAll,omitempty"`
	Secret string `json:"secret,omitempty"`
	AtPersonPhoneNumbers []string `json:"atPersonPhoneNumbers,omitempty"`
	AtPersonList map[string]string `json:"atPersonList,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSNSDingTalkEndpointParam CreateSNSDingTalkEndpoint request param
type CreateSNSDingTalkEndpointParam struct {
	BaseParam
	Params CreateSNSDingTalkEndpointDetailParam `json:"params"`
}
