// Copyright (c) ZStack.io, Inc.

package param

// RemoveRemoteCidrsFromIPsecConnectionDetailParam RemoveRemoteCidrsFromIPsecConnection detail param
type RemoveRemoteCidrsFromIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PeerCidrs []string `json:"peerCidrs" validate:"required"`
}

// RemoveRemoteCidrsFromIPsecConnectionParam RemoveRemoteCidrsFromIPsecConnection request param
type RemoveRemoteCidrsFromIPsecConnectionParam struct {
	BaseParam
	Params RemoveRemoteCidrsFromIPsecConnectionDetailParam `json:"params"`
}
