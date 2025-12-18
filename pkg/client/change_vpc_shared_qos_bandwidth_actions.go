// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVpcSharedQosBandwidth changes VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(uuid string, params param.ChangeVpcSharedQosBandwidthParam) (*view.ChangeVpcSharedQosBandwidthEventView, error) {
	resp := view.ChangeVpcSharedQosBandwidthEventView{}
	if err := cli.Put("v1/vips/sharedqos/{sharedQosUuid}/bandwidth/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
