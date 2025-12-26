// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcIPsecLog gets VpcIPsecLog by uuid
func (cli *ZSClient) GetVpcIPsecLog(uuid string) (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.Get("v1/vpc/virtual-routers/ipseclog", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
