// Copyright (c) ZStack.io, Inc.

package param

// RevokeResourceSharingDetailParam RevokeResourceSharing详细参数
type RevokeResourceSharingDetailParam struct {
	rest []string `json:"resourceUuids" validate:"required"` // 必填
	rest bool `json:"toPublic,omitempty"`
	rest []string `json:"accountUuids,omitempty"`
	rest bool `json:"all,omitempty"`
}

// RevokeResourceSharingParam RevokeResourceSharing请求参数
type RevokeResourceSharingParam struct {
	BaseParam
	Params RevokeResourceSharingDetailParam `json:"params"` // 详细参数
}

