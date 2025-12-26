// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIPsecConnection creates IPsecConnection
func (cli *ZSClient) CreateIPsecConnection(params param.CreateIPsecConnectionParam) (*view.CreateIPsecConnectionEventView, error) {
	resp := view.CreateIPsecConnectionEventView{}
	if err := cli.Post("v1/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
