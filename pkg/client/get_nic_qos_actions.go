// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.Get("v1/vm-instances/{uuid}/nic-qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
