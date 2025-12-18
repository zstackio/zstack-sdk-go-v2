// Copyright (c) ZStack.io, Inc.

package param

// AddRemoteCidrsToIPsecConnectionDetailParam AddRemoteCidrsToIPsecConnection detail param
type AddRemoteCidrsToIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PeerCidrs []string `json:"peerCidrs" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddRemoteCidrsToIPsecConnectionParam AddRemoteCidrsToIPsecConnection request param
type AddRemoteCidrsToIPsecConnectionParam struct {
	BaseParam
	Params AddRemoteCidrsToIPsecConnectionDetailParam `json:"params"`
}
