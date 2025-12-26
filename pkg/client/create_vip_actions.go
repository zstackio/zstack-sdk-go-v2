// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVip creates Vip
func (cli *ZSClient) CreateVip(params param.CreateVipParam) (*view.CreateVipEventView, error) {
	resp := view.CreateVipEventView{}
	if err := cli.Post("v1/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
