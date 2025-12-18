// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVipQos gets VipQos by uuid
func (cli *ZSClient) GetVipQos(uuid string) (*view.GetVipQosView, error) {
	var resp view.GetVipQosView
	if err := cli.Get("v1/vip/{uuid}/vip-qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
