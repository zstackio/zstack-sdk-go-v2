// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachL3NetworksToIPsecConnection operates on L3NetworksToIPsecConnection
func (cli *ZSClient) AttachL3NetworksToIPsecConnection(params param.AttachL3NetworksToIPsecConnectionParam) (*view.AttachL3NetworksToIPsecConnectionEventView, error) {
	resp := view.AttachL3NetworksToIPsecConnectionEventView{}
	if err := cli.Post("v1/ipsec/{uuid}/l3networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
