// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddRemoteCidrsToIPsecConnection adds RemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(params param.AddRemoteCidrsToIPsecConnectionParam) (*view.AddRemoteCidrsToIPsecConnectionEventView, error) {
	resp := view.AddRemoteCidrsToIPsecConnectionEventView{}
	if err := cli.Post("v1/ipsec/{uuid}/remote-cidrs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
