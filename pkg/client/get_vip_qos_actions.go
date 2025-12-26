// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVipQos gets VipQos by uuid
func (cli *ZSClient) GetVipQos(uuid string) (*view.GetVipQosView, error) {
	var resp view.GetVipQosView
	if err := cli.Get("v1/vip/{uuid}/vip-qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
