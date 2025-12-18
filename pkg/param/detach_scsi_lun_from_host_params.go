// Copyright (c) ZStack.io, Inc.

package param

// DetachScsiLunFromHostDetailParam DetachScsiLunFromHost detail param
type DetachScsiLunFromHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// DetachScsiLunFromHostParam DetachScsiLunFromHost request param
type DetachScsiLunFromHostParam struct {
	BaseParam
	Params DetachScsiLunFromHostDetailParam `json:"params"`
}
