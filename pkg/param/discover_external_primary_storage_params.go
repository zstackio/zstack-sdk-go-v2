// Copyright (c) ZStack.io, Inc.

package param

// DiscoverExternalPrimaryStorageDetailParam DiscoverExternalPrimaryStorage detail param
type DiscoverExternalPrimaryStorageDetailParam struct {
	Url string `json:"url" validate:"required"`
	Identity string `json:"identity,omitempty"`
	Config string `json:"config,omitempty"`
}

// DiscoverExternalPrimaryStorageParam DiscoverExternalPrimaryStorage request param
type DiscoverExternalPrimaryStorageParam struct {
	BaseParam
	Params DiscoverExternalPrimaryStorageDetailParam `json:"params"`
}
