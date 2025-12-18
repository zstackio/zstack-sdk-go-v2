// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVniRange updates VniRange
func (cli *ZSClient) UpdateVniRange(uuid string, params param.UpdateVniRangeParam) (*view.UpdateVniRangeEventView, error) {
	resp := view.UpdateVniRangeEventView{}
	if err := cli.Put("v1/l2-networks/vxlan-pool/vni-ranges/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
