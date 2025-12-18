// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVolumeQos 获取VolumeQos详情
func (cli *ZSClient) GetVolumeQos(uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.Get("v1/volumes/{uuid}/qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

