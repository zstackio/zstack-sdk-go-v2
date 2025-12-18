// Copyright (c) ZStack.io, Inc.

package param

// RevokeResourceSharingDetailParam RevokeResourceSharing detail param
type RevokeResourceSharingDetailParam struct {
	ResourceUuids []string `json:"resourceUuids" validate:"required"`
	ToPublic bool `json:"toPublic,omitempty"`
	AccountUuids []string `json:"accountUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// RevokeResourceSharingParam RevokeResourceSharing request param
type RevokeResourceSharingParam struct {
	BaseParam
	Params RevokeResourceSharingDetailParam `json:"params"`
}
