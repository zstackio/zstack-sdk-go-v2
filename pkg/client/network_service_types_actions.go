// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetNetworkServiceTypes 获取NetworkServiceTypes详情
func (cli *ZSClient) GetNetworkServiceTypes(uuid string) (*view.GetNetworkServiceTypesView, error) {
	var resp view.GetNetworkServiceTypesView
	if err := cli.Get("v1/network-services/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

