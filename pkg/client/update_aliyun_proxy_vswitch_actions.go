// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunProxyVSwitch updates AliyunProxyVSwitch
func (cli *ZSClient) UpdateAliyunProxyVSwitch(uuid string, params param.UpdateAliyunProxyVSwitchParam) (*view.UpdateAliyunProxyVSwitchEventView, error) {
	resp := view.UpdateAliyunProxyVSwitchEventView{}
	if err := cli.Put("v1/aliyun-proxy/vswitches/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
