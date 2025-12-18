// Copyright (c) ZStack.io, Inc.

package param

// SNSWeComTestConnectionDetailParam SNSWeComTestConnection detail param
type SNSWeComTestConnectionDetailParam struct {
	Url string `json:"url,omitempty"`
	AtAll bool `json:"atAll,omitempty"`
	AtPersonUserIds []string `json:"atPersonUserIds,omitempty"`
	TestMsg string `json:"testMsg" validate:"required"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// SNSWeComTestConnectionParam SNSWeComTestConnection request param
type SNSWeComTestConnectionParam struct {
	BaseParam
	Params SNSWeComTestConnectionDetailParam `json:"params"`
}
