// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunProxyVSwitch creates AliyunProxyVSwitch
func (cli *ZSClient) CreateAliyunProxyVSwitch(params param.CreateAliyunProxyVSwitchParam) (*view.CreateAliyunProxyVSwitchEventView, error) {
	resp := view.CreateAliyunProxyVSwitchEventView{}
	if err := cli.Post("v1/aliyun-proxy/vpcs/vswitches", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
