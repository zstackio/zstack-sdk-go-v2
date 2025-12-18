// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachVipToVpcSharedQos 操作VipToVpcSharedQos
func (cli *ZSClient) AttachVipToVpcSharedQos(params param.AttachVipToVpcSharedQosParam) (*view.AttachVipToVpcSharedQosEventView, error) {
	resp := view.AttachVipToVpcSharedQosEventView{}
	if err := cli.Post("v1/vips/sharedqos/{sharedQosUuid}/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

