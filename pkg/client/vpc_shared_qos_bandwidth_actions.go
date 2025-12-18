// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVpcSharedQosBandwidth 操作VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(params param.ChangeVpcSharedQosBandwidthParam) (*view.ChangeVpcSharedQosBandwidthEventView, error) {
	resp := view.ChangeVpcSharedQosBandwidthEventView{}
	if err := cli.Post("v1/vips/sharedqos/{sharedQosUuid}/bandwidth/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

