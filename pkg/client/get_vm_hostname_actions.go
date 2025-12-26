// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmHostname gets VmHostname by uuid
func (cli *ZSClient) GetVmHostname(uuid string) (*view.GetVmHostnameView, error) {
	var resp view.GetVmHostnameView
	if err := cli.Get("v1/vm-instances/{uuid}/hostnames", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
