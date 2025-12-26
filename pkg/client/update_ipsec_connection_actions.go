// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIPsecConnection updates IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(uuid string, params param.UpdateIPsecConnectionParam) (*view.UpdateIPsecConnectionEventView, error) {
	resp := view.UpdateIPsecConnectionEventView{}
	if err := cli.Put("v1/ipsec/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
