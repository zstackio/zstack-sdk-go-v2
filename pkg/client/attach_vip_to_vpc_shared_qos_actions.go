// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachVipToVpcSharedQos operates on VipToVpcSharedQos
func (cli *ZSClient) AttachVipToVpcSharedQos(params param.AttachVipToVpcSharedQosParam) (*view.AttachVipToVpcSharedQosEventView, error) {
	resp := view.AttachVipToVpcSharedQosEventView{}
	if err := cli.Post("v1/vips/sharedqos/{sharedQosUuid}/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
