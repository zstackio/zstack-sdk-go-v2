// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcIPsecLog 获取VpcIPsecLog详情
func (cli *ZSClient) GetVpcIPsecLog(uuid string) (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.Get("v1/vpc/virtual-routers/ipseclog", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

