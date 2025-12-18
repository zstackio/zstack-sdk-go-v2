// Copyright (c) ZStack.io, Inc.

package param

// IsVfNicAvailableInL3NetworkDetailParam IsVfNicAvailableInL3Network detail param
type IsVfNicAvailableInL3NetworkDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
}

// IsVfNicAvailableInL3NetworkParam IsVfNicAvailableInL3Network request param
type IsVfNicAvailableInL3NetworkParam struct {
	BaseParam
	Params IsVfNicAvailableInL3NetworkDetailParam `json:"params"`
}
