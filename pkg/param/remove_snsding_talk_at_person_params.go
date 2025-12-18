// Copyright (c) ZStack.io, Inc.

package param

// RemoveSNSDingTalkAtPersonDetailParam RemoveSNSDingTalkAtPerson详细参数
type RemoveSNSDingTalkAtPersonDetailParam struct {
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"phoneNumber" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// RemoveSNSDingTalkAtPersonParam RemoveSNSDingTalkAtPerson请求参数
type RemoveSNSDingTalkAtPersonParam struct {
	BaseParam
	Params RemoveSNSDingTalkAtPersonDetailParam `json:"params"` // 详细参数
}

