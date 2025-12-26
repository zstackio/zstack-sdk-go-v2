// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVniRange updates VniRange
func (cli *ZSClient) UpdateVniRange(uuid string, params param.UpdateVniRangeParam) (*view.UpdateVniRangeEventView, error) {
	resp := view.UpdateVniRangeEventView{}
	if err := cli.Put("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
