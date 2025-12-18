// Copyright (c) ZStack.io, Inc.

package param

// SyncIdentityFromRemoteDetailParam SyncIdentityFromRemote detail param
type SyncIdentityFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SyncIdentityFromRemoteParam SyncIdentityFromRemote request param
type SyncIdentityFromRemoteParam struct {
	BaseParam
	Params SyncIdentityFromRemoteDetailParam `json:"params"`
}
