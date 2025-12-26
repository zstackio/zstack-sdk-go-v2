// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIpRange updates IpRange
func (cli *ZSClient) UpdateIpRange(uuid string, params param.UpdateIpRangeParam) (*view.UpdateIpRangeEventView, error) {
	resp := view.UpdateIpRangeEventView{}
	if err := cli.Put("v1/l3-networks/ip-ranges/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
