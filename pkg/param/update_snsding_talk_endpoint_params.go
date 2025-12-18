// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSDingTalkEndpointDetailParam UpdateSNSDingTalkEndpoint detail param
type UpdateSNSDingTalkEndpointDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	Secret string `json:"secret,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
}

// UpdateSNSDingTalkEndpointParam UpdateSNSDingTalkEndpoint request param
type UpdateSNSDingTalkEndpointParam struct {
	BaseParam
	Params UpdateSNSDingTalkEndpointDetailParam `json:"params"`
}
