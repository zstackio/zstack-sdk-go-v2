// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIPsecConnection changes IPsecConnection
func (cli *ZSClient) ChangeIPsecConnection(uuid string, params param.ChangeIPsecConnectionParam) (*view.ChangeIPsecConnectionEventView, error) {
	resp := view.ChangeIPsecConnectionEventView{}
	if err := cli.Put("v1/ipsec/config/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
