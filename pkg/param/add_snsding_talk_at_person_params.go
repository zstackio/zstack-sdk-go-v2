// Copyright (c) ZStack.io, Inc.

package param

// AddSNSDingTalkAtPersonDetailParam AddSNSDingTalkAtPerson详细参数
type AddSNSDingTalkAtPersonDetailParam struct {
	rest string `json:"phoneNumber" validate:"required"` // 必填
	rest string `json:"endpointUuid" validate:"required"` // 必填
	rest string `json:"remark,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddSNSDingTalkAtPersonParam AddSNSDingTalkAtPerson请求参数
type AddSNSDingTalkAtPersonParam struct {
	BaseParam
	Params AddSNSDingTalkAtPersonDetailParam `json:"params"` // 详细参数
}

