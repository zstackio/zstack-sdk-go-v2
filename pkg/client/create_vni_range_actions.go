// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVniRange creates VniRange
func (cli *ZSClient) CreateVniRange(params param.CreateVniRangeParam) (*view.CreateVniRangeEventView, error) {
	resp := view.CreateVniRangeEventView{}
	if err := cli.Post("v1/l2-networks/vxlan-pool/{l2NetworkUuid}/vni-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
