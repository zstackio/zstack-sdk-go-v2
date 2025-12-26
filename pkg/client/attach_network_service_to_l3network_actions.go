// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachNetworkServiceToL3Network operates on NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(params param.AttachNetworkServiceToL3NetworkParam) (*view.AttachNetworkServiceToL3NetworkEventView, error) {
	resp := view.AttachNetworkServiceToL3NetworkEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/network-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
