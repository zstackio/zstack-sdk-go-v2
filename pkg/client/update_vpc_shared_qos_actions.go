// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVpcSharedQos updates VpcSharedQos
func (cli *ZSClient) UpdateVpcSharedQos(uuid string, params param.UpdateVpcSharedQosParam) (*view.UpdateVpcSharedQosEventView, error) {
	resp := view.UpdateVpcSharedQosEventView{}
	if err := cli.Put("v1/vips/sharedqos/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
