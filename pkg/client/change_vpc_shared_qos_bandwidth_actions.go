// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVpcSharedQosBandwidth changes VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(uuid string, params param.ChangeVpcSharedQosBandwidthParam) (*view.ChangeVpcSharedQosBandwidthEventView, error) {
	resp := view.ChangeVpcSharedQosBandwidthEventView{}
	if err := cli.Put("v1/vips/sharedqos/{sharedQosUuid}/bandwidth/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
