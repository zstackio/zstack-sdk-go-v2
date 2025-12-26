// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVCenter adds VCenter
func (cli *ZSClient) AddVCenter(params param.AddVCenterParam) (*view.AddVCenterEventView, error) {
	resp := view.AddVCenterEventView{}
	if err := cli.Post("v1/vcenters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
