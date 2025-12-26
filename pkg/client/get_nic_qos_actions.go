// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.Get("v1/vm-instances/{uuid}/nic-qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
