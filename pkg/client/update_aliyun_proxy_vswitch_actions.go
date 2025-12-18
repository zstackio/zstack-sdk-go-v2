// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunProxyVSwitch updates AliyunProxyVSwitch
func (cli *ZSClient) UpdateAliyunProxyVSwitch(uuid string, params param.UpdateAliyunProxyVSwitchParam) (*view.UpdateAliyunProxyVSwitchEventView, error) {
	resp := view.UpdateAliyunProxyVSwitchEventView{}
	if err := cli.Put("v1/aliyun-proxy/vswitches/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
